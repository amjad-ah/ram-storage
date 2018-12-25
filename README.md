# Simple stupid key-value store
> a very simple key-value store on RAM

# How?

> install the package 

> `go run *.go -port :8080`

# HTTP Methods:

> POST `/` with: 
```
request:
{
    "name": "Amjad",
    "age": 24
}
response:
{
    "data": {
        "name": "Amjad",
        "age": 24
    },
    "success": true
}
```

> GET `/` will return all keys and values that had been stored:
```
{
    "data": {
        "name": "Amjad",
        "age": 24
    },
    "success": true
}
```

> DELETE `/:key` will delete a specific key for example `DELETE /name`:
```
{
    "data": {
        "age": 24
    },
    "success": true
}
```