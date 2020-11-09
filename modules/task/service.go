package task

import (
    "fmt"
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

func (service *Service) CreateTask(owner string, data TaskCreate) (*entities.Task, error) {
    task := &entities.Task{Name: data.Name, Detail: data.Detail, Owner: owner}
    err := service.repo.CreateTask(task)
    return task, err
}

func (service *Service) GetAllTasks(owner string) (*[]entities.Task, error) {
    return service.repo.GetAllTasks(owner)
}

func (service *Service) GetDetailTask(task_id int, owner string) (*entities.Task, error) {
    return service.repo.GetDetailTask(task_id, owner)
}

func (service *Service) UpdateTask(task_id int, data TaskUpdate) error {
    if data.Name != nil {
        if *data.Name == "" {
            return fmt.Errorf("Name of the task must be non-empty")
        }
        err := service.repo.UpdateNameTask(task_id, *data.Name)
        if err != nil {
            return err
        }
    }
    if data.Detail != nil {
        err := service.repo.UpdateDetailTask(task_id, *data.Detail)
        if err != nil {
            return err
        }
    }
    return nil
}

func (service *Service) ArchiveTask(task_id int) error {
    return service.repo.UpdateDoneTask(task_id, true)
}


func (service *Service) RestoreTask(task_id int) error {
    return service.repo.UpdateDoneTask(task_id, false)
}

func (service *Service) DeleteTask(task_id int) error {
    task, err := service.repo.GetTaskByID(task_id)
    if err != nil {
        return err
    }
    err = service.repo.DeleteTask(task_id)
    if err != nil {
        return err
    }
    deleted_task := entities.DeletedTask(*task)
    return service.repo.CreateDeletedTask(&deleted_task)
}
