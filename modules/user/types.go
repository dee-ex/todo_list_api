package user

type UserRegister struct {
    Username string `json: "username"`
    Email string `json: "email"`
    Password string `json: "password"`
}

type UserLogin struct {
    Username string `json: "username"`
    Password string `json: "password"`
}

type UserUpdate struct {
    Password *string `json: "password"`
}

type TokenLogin struct {
    AccessToken string `json: "access_token"`
}

type ResetPassword struct {
    Email string `json: "email"`
}

type ResetPasswordSubmit struct {
    Password string `json: "password"`
}
