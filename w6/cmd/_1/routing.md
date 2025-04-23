Анатомия веб-сервиса
* Роутинг

```go
/users/123
    -> func GetUser(...) { ... }

    router.HandleFunc(
        "/users/{id:[0-9]+}",
        GetUser,
    )
```