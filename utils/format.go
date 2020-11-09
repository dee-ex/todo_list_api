package utils

import (
    "encoding/json"
    "net/http"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
    response, err := json.Marshal(payload)
    if err != nil {
        HandleError(w, http.StatusInternalServerError, err)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write([]byte(response))
}
