package entities

type User struct {
    ID int `json: "id" gorm: "column:ID"`
    Username string `json: "username" gorm: "column:Username"`
    Email string `json: "email" gorm:  "column:Email"`
    Password string `json: "password" gorm: "column:Password"`
    Deleted bool `json: "deleted" gorm: "column:Deleted"`
}

type Token struct {
    ID int `json: "id" gorm: "column:ID"`
    User string `json: "user" gorm: "column:User"`
    AccessToken string `json: "access_token" gorm: "column:Access_Token"`
}
