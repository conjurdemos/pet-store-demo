package main

import (
    "database/sql"
)

type Store interface {
    CreatePet(pet *Pet) error
    GetPets() ([]*Pet, error)
}

type dbStore struct { // implements `Store` interface
    db *sql.DB // DB connection
}

func (store *dbStore) CreatePet(pet *Pet) error {
    _, err := store.db.Query("INSERT INTO pets(name) VALUES ($1)", pet.Name)

    return err
}

func (store *dbStore) GetPets() ([]*Pet, error) {
    rows, err := store.db.Query("SELECT name from pets")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    pets := []*Pet{}
    for rows.Next() {
        pet := &Pet{}

        if err := rows.Scan(&pet.Name); err != nil {
            return nil, err
        }

        pets = append(pets, pet)
    }
    return pets, nil
}

var store Store

func InitStore(s Store) {
    store = s
}
