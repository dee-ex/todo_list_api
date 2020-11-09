package entities

type Task struct {
    ID int `json: "id" gorm: "column:ID"`
    Name string `json: "name" gorm: "column:Name"`
    Detail string `json: "detail" gorm: "column:Detail"`
    Done bool `json: "done" gorm: "column:Done"`
    Owner string `json: "onwer" gorm: "column:Owner"`
}

type DeletedTask struct {
    ID int `json: "id" gorm: "column:ID"`
    Name string `json: "name" gorm: "column:Name"`
    Detail string `json: "detail" gorm: "column:Detail"`
    Done bool `json: "done" gorm: "column:Done"`
    Owner string `json: "onwer" gorm: "column:Owner"`
}
