package middlewares

import (
    "fmt"
    "context"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/dgrijalva/jwt-go"
    "github.com/dee-ex/todo_list_api/utils"
    "github.com/dee-ex/todo_list_api/entities"
    "github.com/dee-ex/todo_list_api/infrastructures"
)

func ResetpwMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        resetpw_token, err := ValidateToken(vars["resetpw_token"])
        if err != nil {
            utils.HandleError(w, http.StatusUnauthorized, err)
            return
        }
        resetpw_detail, err := ExtractResetpwClaims(resetpw_token)
        if err != nil {
            utils.HandleError(w, http.StatusUnauthorized, err)
            return
        }
        query_resetpw_token, err := GetResetpwTokenByUsername(resetpw_detail.Username)
        if err != nil {
            utils.HandleError(w, http.StatusUnprocessableEntity, err)
            return
        }
        if query_resetpw_token == "" {
            utils.HandleError(w, http.StatusUnauthorized, fmt.Errorf("Username not found or User's token hasn't been created yet"))
            return
        }
        if query_resetpw_token != vars["resetpw_token"] {
            utils.HandleError(w, http.StatusUnauthorized, fmt.Errorf("Token may be old"))
            return
        }
        ctx := context.WithValue(r.Context(), "username", resetpw_detail.Username)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}


func ExtractResetpwClaims(token *jwt.Token) (*ResetpwDetail, error) {
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, fmt.Errorf("Cannot extract data from token")
    }
    username, ok := claims["username"].(string)
    if !ok {
        return nil, fmt.Errorf("Cannot extract username from token")
    }
    return &ResetpwDetail{Username: username}, nil
}


func GetResetpwTokenByUsername(username string) (string, error) {
    db, err := infrastructures.NewMysqlSession()
    if err != nil {
        return "", err
    }
    var resetpw_token string
    res := db.Model(&entities.Token{}).Select("Resetpw_Token").Where("User = ?", username).Find(&resetpw_token)
    return resetpw_token, res.Error
}
