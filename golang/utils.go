package main

import (
    "encoding/json"
    "net/http"
    "fmt"
)

func prettyPrint(payload interface{}) ([]byte) {
    prettyBytes, _ := json.MarshalIndent(payload,"","  ")
    prettyBytes = append(prettyBytes, '\n')

    return prettyBytes
}

func respondWithError(w http.ResponseWriter, err error) {
    fmt.Println(fmt.Errorf("Error: %v", err))
    w.WriteHeader(http.StatusInternalServerError)
    w.Write(prettyPrint(ErrorResponse{
        Error:   err.Error(),
    }))
}
