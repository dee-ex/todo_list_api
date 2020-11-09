package middlewares

import (
    "os"
    "fmt"
    "github.com/dgrijalva/jwt-go"
)

func ValidateToken(token_string string) (*jwt.Token, error) {
    token, err := VerifyToken(token_string)
    if err != nil {
        return nil, fmt.Errorf("Token is unverified")
    }
    if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
        return nil, fmt.Errorf("Token is invalid")
    }
    return token, nil
}

func VerifyToken(token_string string) (*jwt.Token, error) {
    token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(os.Getenv("SECRET_KEY")), nil
    })
    if err != nil {
        return nil, err
    }
    return token, nil
}
