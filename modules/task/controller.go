package task

import (
    "fmt"
    "strings"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/dee-ex/todo_list_api/utils"
    "github.com/dee-ex/todo_list_api/entities"
)

func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
    owner := r.Context().Value("username").(string)
    var data TaskCreate
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        utils.HandleError(w, http.StatusBadRequest, err)
        return
    }
    if data.Name == "" {
        utils.HandleError(w, http.StatusBadRequest, fmt.Errorf("Name of the task must be non-empty"))
        return
    }
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    task, err := service.CreateTask(owner, data)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    utils.RespondJSON(w, http.StatusOK, task)
}

func HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
    owner := r.Context().Value("username").(string)
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    tasks, err := service.GetAllTasks(owner)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    completed, ok := r.URL.Query()["completed"]
    if !ok || strings.ToLower(completed[0]) != "true" {
        utils.RespondJSON(w, http.StatusOK, tasks)
        return
    }
    var completed_tasks []entities.Task
    for _, task := range *tasks {
        if task.Done {
            completed_tasks = append(completed_tasks, task)
        }
    }
    utils.RespondJSON(w, http.StatusOK, completed_tasks)
}

func HandleGetDetailTask(w http.ResponseWriter, r *http.Request) {
    owner := r.Context().Value("username").(string)
    vars := mux.Vars(r)
    task_id, err := strconv.Atoi(vars["task_id"])
    if err != nil {
        utils.HandleError(w, http.StatusNotFound, err)
        return
    }
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    task, err := service.GetDetailTask(task_id, owner)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    // no task has 0 id
    if task.ID == 0 {
        utils.HandleError(w, http.StatusNotFound, fmt.Errorf("You do not own this task"))
        return
    }
    utils.RespondJSON(w, http.StatusOK, task)
}

func HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
    owner := r.Context().Value("username").(string)
    vars := mux.Vars(r)
    task_id, err := strconv.Atoi(vars["task_id"])
    if err != nil {
        utils.HandleError(w, http.StatusNotFound, err)
        return
    }
    var data TaskUpdate
    err = json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        utils.HandleError(w, http.StatusBadRequest, err)
        return
    }
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    task, err := service.GetDetailTask(task_id, owner)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    // no task has 0 id
    if task.ID == 0 {
        utils.HandleError(w, http.StatusNotFound, fmt.Errorf("You do not own this task"))
        return
    }
    err = service.UpdateTask(task_id, data)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    fmt.Fprintf(w, "Updated task successfully")
}

func HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
    owner := r.Context().Value("username").(string)
    vars := mux.Vars(r)
    task_id, err := strconv.Atoi(vars["task_id"])
    if err != nil {
        utils.HandleError(w, http.StatusNotFound, err)
        return
    }
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    task, err := service.GetDetailTask(task_id, owner)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    // no task has 0 id
    if task.ID == 0 {
        utils.HandleError(w, http.StatusNotFound, fmt.Errorf("You do not own this task"))
        return
    }
    err = service.DeleteTask(task_id)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    fmt.Fprintf(w, "Deleted task successfully")
}


func HandleArchiveTask(w http.ResponseWriter, r *http.Request) {
    owner := r.Context().Value("username").(string)
    vars := mux.Vars(r)
    task_id, err := strconv.Atoi(vars["task_id"])
    if err != nil {
        utils.HandleError(w, http.StatusNotFound, err)
        return
    }
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    task, err := service.GetDetailTask(task_id, owner)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    // no task has 0 id
    if task.ID == 0 {
        utils.HandleError(w, http.StatusNotFound, fmt.Errorf("You do not own this task"))
        return
    }
    err = service.ArchiveTask(task_id)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    fmt.Fprintf(w, "Congratulations on archiving task")
}

func HandleRestoreTask(w http.ResponseWriter, r *http.Request) {
    owner := r.Context().Value("username").(string)
    vars := mux.Vars(r)
    task_id, err := strconv.Atoi(vars["task_id"])
    if err != nil {
        utils.HandleError(w, http.StatusNotFound, err)
        return
    }
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    task, err := service.GetDetailTask(task_id, owner)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    // no task has 0 id
    if task.ID == 0 {
        utils.HandleError(w, http.StatusNotFound, fmt.Errorf("You do not own this task"))
        return
    }
    err = service.RestoreTask(task_id)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    fmt.Fprintf(w, "Task has just been restored")
}
