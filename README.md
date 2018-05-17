# pet-store
A demo application creating using the Spring Framework. This application requires access to a database.

### Building
Run `./bin/build`. Requires Docker.

### Running
When running the pet-store, the following environment variables are expected:  
`DB_URL` Url or connection string  
`DB_USERNAME` Username to connect as (not required for secretless)  
`DB_PASSWORD` Password to connect as (not required for secretless)  

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