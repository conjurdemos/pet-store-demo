package golang

import (
    "fmt"
    "net/http"
    "database/sql"
    "os"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

func newRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/hello", handler).Methods("GET")

    r.HandleFunc("/pets", getPetsHandler).Methods("GET")
    r.HandleFunc("/pet", createPetHandler).Methods("POST")
    return r
}

func main() {
    fmt.Println("Starting server...")
    connString := os.Getenv("DB_URL")

    db, err := sql.Open("postgres", connString)
    if err != nil {
        panic(err)
    }
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

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}
