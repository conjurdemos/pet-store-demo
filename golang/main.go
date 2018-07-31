package main

import (
    "fmt"
    "net/http"
    "database/sql"
    "os"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
    "strings"
)

func newRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/vulnerable", vulnerableHandler).Methods("GET")

    r.HandleFunc("/pets", getPetsHandler).Methods("GET")
    r.HandleFunc("/pet", createPetHandler).Methods("POST")
    r.HandleFunc("/pet/{id}", getPetHandler).Methods("GET")
    return r
}

func main() {
    fmt.Println("Starting server...")
    connString := os.Getenv("DB_URL")

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

    r := newRouter()
    fmt.Println("Serving on port 8080")
    http.NewServeMux()
    http.ListenAndServe(":8080", r)
}

func vulnerableHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, strings.Join(os.Environ(), "\n"))
    fmt.Fprintf(w, "\n")
}
