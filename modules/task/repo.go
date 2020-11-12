package task

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

func (repo *Repo) CreateTask(task *entities.Task) error {
    res := repo.db.Create(task)
    return res.Error
}

func (repo *Repo) GetAllTasks(owner string) (*[]entities.Task, error) {
    var tasks []entities.Task
    res := repo.db.Where("Owner = ? AND Deleted = false", owner).Find(&tasks)
    return &tasks, res.Error
}


func (repo *Repo) GetDetailTask(task_id int, owner string) (*entities.Task, error) {
    var task entities.Task
    res := repo.db.Where("ID = ? AND Owner = ? AND Deleted = false", task_id, owner).Find(&task)
    return &task, res.Error
}

func (repo *Repo) GetTaskByID(task_id int) (*entities.Task, error) {
    var task entities.Task
    res := repo.db.Where("ID = ?", task_id).Find(&task)
    return &task, res.Error
}

func (repo *Repo) UpdateNameTask(task_id int, name string) error {
    res := repo.db.Model(&entities.Task{}).Where("ID = ? AND Deleted = false", task_id).Update("Name", name)
    return res.Error
}


func (repo *Repo) UpdateDetailTask(task_id int, detail string) error {
    res := repo.db.Model(&entities.Task{}).Where("ID = ? AND Deleted = false", task_id).Update("Detail", detail)
    return res.Error
}

func (repo *Repo) UpdateDoneTask(task_id int, done bool) error {
    res := repo.db.Model(&entities.Task{}).Where("ID = ? AND Deleted = false", task_id).Update("Done", done)
    return res.Error
}

func (repo *Repo) DeleteTask(task_id int) error {
    res := repo.db.Model(&entities.Task{}).Where("ID = ? AND Deleted = false", task_id).Update("Deleted", true)
    return res.Error
}
