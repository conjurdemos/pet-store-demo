package main

import (
    "fmt"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
    "strings"
    "os"
)

func newRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/vulnerable", vulnerableHandler).Methods("GET")

    r.HandleFunc("/pets", getPetsHandler).Methods("GET")
    r.HandleFunc("/pet", createPetHandler).Methods("POST")
    r.HandleFunc("/pet/{id}", getPetHandler).Methods("GET")
    return r
}

func vulnerableHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, strings.Join(os.Environ(), "\n"))
    fmt.Fprintf(w, "\n")
}

func main() {
    fmt.Println("Starting server...")
    connString, err := LoadDBConfig()
    if err != nil {
        panic(err)
    }

    db, err := sql.Open("postgres", connString)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    InitStore(&dbStore{db: db})

    fmt.Println("Serving on port " + getServerPort())

    r := newRouter()
    http.NewServeMux()
    http.ListenAndServe(":" + getServerPort(), r)
}