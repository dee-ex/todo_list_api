package user

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/dee-ex/todo_list_api/utils"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
    var data UserRegister
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        utils.HandleError(w, http.StatusBadRequest, err)
        return
    }
    // equalize len(x) == 0 || len(y) == 0 || len(z) == 0.
    // make sure no field is empty.
    if len(data.Username)*len(data.Email)*len(data.Password) == 0 {
        utils.HandleError(w, http.StatusBadRequest, fmt.Errorf("Missing some parameters"))
        return
    }
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    user, err := service.CreateUser(data)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    utils.RespondJSON(w, http.StatusOK, user)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
    var data UserLogin
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        utils.HandleError(w, http.StatusBadRequest, err)
        return
    }
    // equalize len(x) == 0 || len(y) == 0 || len(z) == 0.
    // make sure no field is empty.
    if len(data.Username)*len(data.Password) == 0 {
        utils.HandleError(w, http.StatusBadRequest, fmt.Errorf("Missing some parameters"))
        return
    }
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    access_token, err := service.Login(data)
    if err != nil {
        utils.HandleError(w, http.StatusUnauthorized, err)
        return
    }
    // prepare for fetching token
    token := TokenLogin{AccessToken: access_token}
    utils.RespondJSON(w, http.StatusOK, token)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
    username := r.Context().Value("username").(string)
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    err := service.Logout(username)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    fmt.Fprintf(w, "Logged out successfully")
}

func HandleMe(w http.ResponseWriter, r *http.Request) {
    username := r.Context().Value("username").(string)
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    user, err := service.GetUserByUsername(username)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    utils.RespondJSON(w, http.StatusOK, user)
}

func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
    username := r.Context().Value("username").(string)
    var data UserUpdate
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        utils.HandleError(w, http.StatusBadRequest, err)
        return
    }
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    err = service.UpdateUser(username, data)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    fmt.Fprintf(w, "Updated user successfully")
}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
    username := r.Context().Value("username").(string)
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    err := service.DeleteUser(username)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    fmt.Fprintf(w, "Deleted user successfully")
}

func HandleGeneratePath(w http.ResponseWriter, r *http.Request) {
    var data ResetPassword
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        utils.HandleError(w, http.StatusBadRequest, err)
        return
    }
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    path, err := service.GenereatePath(data)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    // send this link to email address
    link := r.Host + r.URL.Path + "/" + path
    fmt.Fprintf(w, link)
}

func HandleResetPassword(w http.ResponseWriter, r *http.Request) {
    username := r.Context().Value("username").(string)
    var data ResetPasswordSubmit
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        utils.HandleError(w, http.StatusBadRequest, err)
        return
    }
    if data.Password == "" {
        utils.HandleError(w, http.StatusBadRequest, fmt.Errorf("Missing some parameters"))
        return
    }
    service := NewService()
    if service == nil {
        utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("Cannot create a session to database"))
        return
    }
    err = service.ResetPassword(username, data)
    if err != nil {
        utils.HandleError(w, http.StatusInternalServerError, err)
        return
    }
    fmt.Fprintf(w, "Password has been reset")
}
