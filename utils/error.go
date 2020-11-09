package utils

import "fmt"
import "net/http"

func HandleError(w http.ResponseWriter, status int, err error) {
    w.WriteHeader(status)
    switch status {
    case 400:
        fmt.Fprintf(w, "400 BAD REQUEST - %s", err)
    case 401:
        fmt.Fprintf(w, "401 UNAUTHORIZED - %s", err)
    case 404:
        fmt.Fprintf(w, "404 NOT FOUND - %s", err)
    case 406:
        fmt.Fprintf(w, "406 NOT ACCEPTABLE - %s", err)
    case 409:
        fmt.Fprintf(w, "409 CONFLICT - %s", err)
    case 422:
        fmt.Fprintf(w, "422 UNPROCESSABLE ENTITY - %s", err)
    case 500:
        fmt.Fprintf(w, "500 INTERNAL SERVER ERROR - %s", err)
    default:
        fmt.Fprintf(w, "ERROR - %s", err)
    }
}
