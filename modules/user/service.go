package user

import (
    "os"
    "fmt"
    "time"
    "strings"
    "regexp"
    "github.com/dgrijalva/jwt-go"
    "github.com/dee-ex/todo_list_api/entities"
)

type Service struct {
    repo *Repo
}

func NewService() *Service {
    repo := NewRepo()
    if repo == nil {
        return nil
    }
    return &Service{repo: repo}
}

func (service *Service) CreateUser(data UserRegister) (*entities.User, error) {
    err := service.ValidateUsername(data.Username)
    if err != nil {
        return nil, err
    }
    err = service.ValidateEmail(data.Email)
    if err != nil {
        return nil, err
    }
    user := &entities.User{Username: data.Username, Email: data.Email, Password: data.Password}
    err = service.repo.CreateUser(user)
    if err != nil {
        return nil, err
    }
    token := &entities.Token{User: user.Username}
    err = service.repo.CreateToken(token)
    return user, err
}

func (service *Service) ValidateUsername(username string) error {
    // make sure no space in username
    if strings.Contains(username, " ") {
        return fmt.Errorf("Username cannot have space(s)")
    }
    query_username, err := service.repo.GetUsername(username)
    if err != nil {
        return err
    }
    // make sure username does not exists
    if query_username != "" {
        return fmt.Errorf("Username has already been taken")
    }
    return nil
}

func (service *Service) ValidateEmail(email string) error {
    re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    // check email format by regex
    if !re.MatchString(email) {
        return fmt.Errorf("Email is invalid")
    }
    query_email, err := service.repo.GetEmail(email)
    if err != nil {
        return err
    }
    // make sure email does not exists
    if query_email != "" {
        return fmt.Errorf("Email has already been taken")
    }
    return nil
}

func (service *Service) Login(data UserLogin) (string, error) {
    query_username, err := service.repo.GetUsername(data.Username)
    if err != nil {
        return "", err
    }
    if query_username == "" {
        return "", fmt.Errorf("Username not found")
    }
    query_user, err := service.repo.GetUserByUsernameAndPassword(data.Username, data.Password)
    if err != nil {
        return "", err
    }
    // no user has 0 id
    // => 0 id == username & password does not mach each other
    if query_user.ID == 0 {
        return "", fmt.Errorf("Password is wrong")
    }
    access_token, err := CreateAccessToken(query_user.Username)
    if err != nil {
        return "", err
    }
    // update token to database
    err = service.repo.UpdateAccessToken(query_user.Username, access_token)
    return access_token, err
}

func (service *Service) Logout(username string) error {
    // clear the access token
    return service.repo.UpdateAccessToken(username, "")
}

func (service *Service) GetUserByUsername(username string) (*entities.User, error) {
    return service.repo.GetUserByUsername(username)
}

func (service *Service) UpdateUser(username string, data UserUpdate) error {
    if data.Password != nil {
        err := service.repo.UpdatePassword(username, *data.Password)
        return err
    }
    return nil
}

func (service *Service) DeleteUser(username string) error {
    user, err := service.repo.GetUserByUsername(username)
    if err != nil {
        return err
    }
    err = service.repo.DeleteUser(username)
    if err != nil {
        return err
    }
    deleted_user := entities.DeletedUser(*user)
    err = service.repo.CreateDeletedUser(&deleted_user)
    if err != nil {
        return err
    }
    // clear the access token
    return service.repo.UpdateAccessToken(username, "")
}

func (service *Service) GenereatePath(data ResetPassword) (string, error) {
    query_user, err := service.repo.GetUserByEmail(data.Email)
    if err != nil {
        return "", err
    }
    if query_user.ID == 0 {
        return "", fmt.Errorf("Email not found")
    }
    resetpw_token, err := CreateResetpwToken(query_user.Username)
    if err != nil {
        return "", err
    }
    // update token to database
    err = service.repo.UpdateResetpwToken(query_user.Username, resetpw_token)
    return resetpw_token, err
}

func (service *Service) ResetPassword(username string, data ResetPasswordSubmit) error {
    err := service.repo.UpdatePassword(username, data.Password)
    if err != nil {
        return err
    }
    // clear the resetpw token
    // maka sure it's only used once
    err = service.repo.UpdateResetpwToken(username, "")
    if err != nil {
        return err
    }
    // clear the access token
    return service.repo.UpdateAccessToken(username, "")
}

func CreateAccessToken(username string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    token_claims := token.Claims.(jwt.MapClaims)
    token_claims["authorized"] = true
    token_claims["username"] = username
    token_claims["exp"] = time.Now().Add(60*time.Minute).Unix()
    access_token, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
    if err != nil {
        return "", err
    }
    return access_token, nil
}


func CreateResetpwToken(username string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    token_claims := token.Claims.(jwt.MapClaims)
    token_claims["authorized"] = true
    token_claims["username"] = username
    token_claims["exp"] = time.Now().Add(5*time.Minute).Unix()
    resetpw_token, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
    if err != nil {
        return "", err
    }
    return resetpw_token, nil
}
