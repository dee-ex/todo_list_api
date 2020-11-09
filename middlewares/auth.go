package middlewares

import (
    "fmt"
    "context"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "github.com/dee-ex/todo_list_api/utils"
    "github.com/dee-ex/todo_list_api/entities"
    "github.com/dee-ex/todo_list_api/infrastructures"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        access_token_string := r.Header.Get("access_token")
        access_token, err := ValidateToken(access_token_string)
        if err != nil {
            utils.HandleError(w, http.StatusUnauthorized, err)
            return
        }
        access_detail, err := ExtractAccessClaims(access_token)
        if err != nil {
            utils.HandleError(w, http.StatusUnauthorized, err)
            return
        }
        query_access_token, err := GetAccessTokenByUsername(access_detail.Username)
        if err != nil {
            utils.HandleError(w, http.StatusUnprocessableEntity, err)
            return
        }
        // this can jump into 3 cases.
        // Case 1: access_detail.Username is empty == token didn't contain username
        // Case.2: username not found.
        // Case.3: token of this user hasn't been created yet.
        if query_access_token == "" {
            utils.HandleError(w, http.StatusUnauthorized, fmt.Errorf("Username not found or User's token hasn't been created yet"))
            return
        }
        if query_access_token != access_token_string {
            utils.HandleError(w, http.StatusUnauthorized, fmt.Errorf("Token may be old"))
            return
        }
        ctx := context.WithValue(r.Context(), "username", access_detail.Username)
        //ctx = context.WithValue(ctx, "another_key", another_value)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func ExtractAccessClaims(token *jwt.Token) (*AccessDetail, error) {
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, fmt.Errorf("Cannot extract data from token")
    }
    username, ok := claims["username"].(string)
    if !ok {
        return nil, fmt.Errorf("Cannot extract username from token")
    }
    return &AccessDetail{Username: username}, nil
}

func GetAccessTokenByUsername(username string) (string, error) {
    db, err := infrastructures.NewMysqlSession()
    if err != nil {
        return "", err
    }
    var access_token string
    res := db.Model(&entities.Token{}).Select("Access_Token").Where("User = ?", username).Find(&access_token)
    return access_token, res.Error
}
