package golang

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Pet struct {
    Name     string `json:"name"`
}

type ErrorResponse struct {
   Error string `json:"error"`
}

func errorJSONBytes(err error) []byte {
    errBytes, _ := json.Marshal(ErrorResponse{
        Error:   err.Error(),
    })
    return errBytes
}

func getPetsHandler(w http.ResponseWriter, r *http.Request) {
    pets, err := store.GetPets()
    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(errorJSONBytes(err))
        return
    }

    petListBytes, err := json.Marshal(pets)
    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Write(petListBytes)
}

func createPetHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var pet Pet
    err := decoder.Decode(&pet)
    if err != nil {
        fmt.Println(fmt.Errorf("Error: %v", err))
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(errorJSONBytes(err))
        return
    }

    err = store.CreatePet(&pet)
    if err != nil {
        fmt.Println(err)
    }

    w.WriteHeader(http.StatusCreated)
}
