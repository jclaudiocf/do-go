# Project
> Device manager to practice Golang.

# Technologies
- mux (Restful API)
- gorm (GO ORM)
- dep (Dependency manager)

# Endpoint
- `GET .../device&page=0&size=10` (Get all devices)
- `POST .../device` (Create a new device)
- `PUT or PATCH .../device/{id}` (Update an existing device)
- `DELETE .../device/{id}` (Delete an existing device)

JSON example for request `POST` and `PUT or PATCH`:
```
{
    "nickname": "example",
    "owner": "example",
    "ip" : "example
}
```