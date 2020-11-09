package main

import (
    "os"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/dee-ex/todo_list_api/middlewares"
    "github.com/dee-ex/todo_list_api/modules/user"
    "github.com/dee-ex/todo_list_api/modules/task"
)

func StartServer() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/user/register", user.HandleRegister).Methods("POST")
    router.HandleFunc("/user/login", user.HandleLogin).Methods("POST")
    router.HandleFunc("/user/reset-password", user.HandleGeneratePath).Methods("POST")

    resetpw_router := router.PathPrefix("/user/reset-password/{resetpw_token}").Subrouter()
    resetpw_router.Use(middlewares.ResetpwMiddleware)
    resetpw_router.HandleFunc("", user.HandleResetPassword).Methods("POST")

    logout_router := router.PathPrefix("/user/logout").Subrouter()
    logout_router.Use(middlewares.AuthMiddleware)
    logout_router.HandleFunc("", user.HandleLogout).Methods("POST")

    me_router := router.PathPrefix("/user/me").Subrouter()
    me_router.Use(middlewares.AuthMiddleware)
    me_router.HandleFunc("", user.HandleMe).Methods("GET")
    me_router.HandleFunc("", user.HandleUpdateUser).Methods("PUT")
    me_router.HandleFunc("", user.HandleDeleteUser).Methods("DELETE")

    task_router := router.PathPrefix("/task").Subrouter()
    task_router.Use(middlewares.AuthMiddleware)
    task_router.HandleFunc("", task.HandleCreateTask).Methods("POST")
    task_router.HandleFunc("", task.HandleGetAllTasks).Methods("GET")
    task_router.HandleFunc("/{task_id:[0-9]+}", task.HandleGetDetailTask).Methods("GET")
    task_router.HandleFunc("/{task_id:[0-9]+}", task.HandleUpdateTask).Methods("PUT")
    task_router.HandleFunc("/{task_id:[0-9]+}", task.HandleDeleteTask).Methods("DELETE")
    task_router.HandleFunc("/{task_id:[0-9]+}/archive", task.HandleArchiveTask).Methods("PUT")
    task_router.HandleFunc("/{task_id:[0-9]+}/archive", task.HandleRestoreTask).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
    log.Println("SERVER STARTS")
    os.Setenv("SECRET_KEY", "1080100020002000")
    StartServer()
}
