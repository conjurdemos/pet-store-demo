package main

import (
    "encoding/json"
    "net/http"
    "fmt"
    "os"
    "net/url"
)

// pretty prints JSON with indentation
func prettyPrint(payload interface{}) ([]byte) {
    prettyBytes, _ := json.MarshalIndent(payload,"","  ")
    prettyBytes = append(prettyBytes, '\n')

    return prettyBytes
}

// writes error to response as JSON payload
func respondWithError(w http.ResponseWriter, err error) {
    fmt.Println(fmt.Errorf("Error: %v", err))
    w.WriteHeader(http.StatusInternalServerError)
    w.Write(prettyPrint(ErrorResponse{
        Error:   err.Error(),
    }))
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

// DefaultPort is the default port to use if one is not specified by the PORT environment variable
const DefaultPort = "8080";
func getServerPort() (string) {
    port := os.Getenv("PORT");
    if port != "" {
        return port;
    }

    return DefaultPort;
}

