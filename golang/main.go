package main

import (
    "fmt"
    "net/http"
    "database/sql"
    "os"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
    "strings"
    "net/url"
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

    r := newRouter()
    fmt.Println("Serving on port 8080")
    http.NewServeMux()
    http.ListenAndServe(":8080", r)
}

func vulnerableHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, strings.Join(os.Environ(), "\n"))
    fmt.Fprintf(w, "\n")
}

// LoadDBConfig parses an URL into options that can be used to connect to PostgreSQL.
func LoadDBConfig() (string, error) {
    connString := os.Getenv("DB_URL")
    parsedUrl, err := url.Parse(connString)
    if err != nil {
        return "", fmt.Errorf(`DB_URL="%s" %s`, connString, err.Error())
    }

    user := parsedUrl.User
    // username and password
    if dbPassword, ok := os.LookupEnv("DB_PASSWORD"); ok {
        user = url.UserPassword(user.Username(), dbPassword)
    }
    if dbUser, ok := os.LookupEnv("DB_USERNAME"); ok {
        uPassword, uPasswordSet := user.Password()
        if uPasswordSet {
            user = url.UserPassword(dbUser, uPassword)
        } else {
            user = url.User(dbUser)
        }

    }

    parsedUrl.User = user;

    return parsedUrl.String(), nil
}
