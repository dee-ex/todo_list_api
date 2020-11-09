package user

import (
    "gorm.io/gorm"
    "github.com/dee-ex/todo_list_api/entities"
    "github.com/dee-ex/todo_list_api/infrastructures"
)

type Repo struct {
    db *gorm.DB
}

func NewRepo() *Repo {
    db, err := infrastructures.NewMysqlSession()
    if err != nil {
        return nil
    }
    return &Repo{db: db}
}

func (repo *Repo) CreateUser(user *entities.User) error {
    res := repo.db.Create(user)
    return res.Error
}

func (repo *Repo) GetUsername(username string) (string, error) {
    var query_username string
    res := repo.db.Model(&entities.User{}).Select("Username").Where("Username = ?", username).Find(&query_username)
    return query_username, res.Error
}


func (repo *Repo) GetEmail(email string) (string, error) {
    var query_email string
    res := repo.db.Model(&entities.User{}).Select("Email").Where("Email = ?", email).Find(&query_email)
    return query_email, res.Error
}

func (repo *Repo) GetUserByUsername (username string) (*entities.User, error) {
    var user entities.User
    res := repo.db.Where("Username = ?", username).Find(&user)
    return &user, res.Error
}

func (repo *Repo) GetUserByEmail (email string) (*entities.User, error) {
    var user entities.User
    res := repo.db.Where("Email = ?", email).Find(&user)
    return &user, res.Error
}

func (repo *Repo) GetUserByUsernameAndPassword(username, password string) (*entities.User, error) {
    var user entities.User
    res := repo.db.Where("Username = ? AND Password = ?", username, password).Find(&user)
    return &user, res.Error
}

func (repo *Repo) UpdatePassword(username, password string) error {
    res := repo.db.Model(&entities.User{}).Where("Username = ?", username).Update("Password", password)
    return res.Error
}

func (repo *Repo) DeleteUser(username string) error {
    res := repo.db.Where("Username = ?", username).Delete(&entities.User{})
    return res.Error
}

func (repo *Repo) CreateToken(token *entities.Token) error {
    res := repo.db.Create(token)
    return res.Error
}

func (repo *Repo) UpdateAccessToken(username, access_token string) error {
    res := repo.db.Model(&entities.Token{}).Where("User = ?", username).Update("Access_Token", access_token)
    return res.Error
}


func (repo *Repo) UpdateResetpwToken(username, resetpw_token string) error {
    res := repo.db.Model(&entities.Token{}).Where("User = ?", username).Update("Resetpw_Token", resetpw_token)
    return res.Error
}

func (repo *Repo) CreateDeletedUser(deleted_user *entities.DeletedUser) error {
    res := repo.db.Create(deleted_user)
    return res.Error
}
