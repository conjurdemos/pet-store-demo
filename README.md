# pet-store
A demo application creating using the Spring Framework. This application requires access to a database.

When running the pet-store, the following environment variables are expected:  
- `DB_URL` Url or connection string.
- `DB_USERNAME` Username to connect as (not required for secretless).
- `DB_PASSWORD` Password to connect as (not required for secretless).
- `DB_PLATFORM` Platform to use in the DDL or DML scripts (such as schema-${platform}.sql
  or data-${platform}.sql). Supported values of `DB_PLATFORM` are `mysql`, `mssql`, and `postgres`.

### Routes
The demo application mocks a pet store service which controls an inventory of pets in a persistent database. The following routes are exposed:

---
`GET` `/pets`  
List all pets in inventory
##### Returns
`200`  
An array of pets in the response body
```
[
  {
    "name": "Scooter"
  },
  {
    "name": "Sparky"
  }
]
```

---
`POST` `/pet`  
Add a pet to the inventory
##### Request
###### Headers
`Content-Type: application/json`
###### Body
```
{
  "name": "Scooter"
}
```
##### Returns
`201`

---
`GET` `/pet/{id}`  
Retrieve information on a pet
##### Returns
`404`
`200`
```
{
  "id": 1
  "name": "Scooter"
}
```
---
`DELETE` `/pet/{id}`  
Remove a pet from inventory
##### Returns
`404`
`200`

---
`GET` `/vulnerable`
Return a JSON representation of all environment variables that
the app knows about
##### Returns
`200`

# Contributing

To learn more about contributing to this repository, please see [CONTRIBUTING.md](CONTRIBUTING.md).

# License

The Pet Store demo app is licensed under Apache License 2.0 - see [`LICENSE.md`](LICENSE.md) for more details.
