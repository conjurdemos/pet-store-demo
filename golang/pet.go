package main

import (
    "encoding/json"
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    "strconv"
)

type Pet struct {
    ID     int64 `json:"id"`
    Name   string `json:"name"`
}

type ErrorResponse struct {
   Error string `json:"error"`
}

// UTILS
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

// HANDLERS
func getPetHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, err)
        return
    }

    pet, err := store.GetPet(int64(id))
    if err != nil {
        respondWithError(w, err)
        return
    }

    w.Write(prettyPrint(pet))
}

func getPetsHandler(w http.ResponseWriter, r *http.Request) {
    pets, err := store.GetPets()
    if err != nil {
        respondWithError(w, err)
        return
    }

    w.Write(prettyPrint(pets))
}

func createPetHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var pet Pet
    err := decoder.Decode(&pet)
    if err != nil {
        respondWithError(w, err)
        return
    }

    err = store.CreatePet(&pet)
    if err != nil {
        respondWithError(w, err)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
