Анатомия веб-сервиса
* Middleware / фреймворки
```go
Site := myServeMux
Site = checkAuth(Handler)
Site = accessLog(Handler)
Site = avoidPanic(Handler)

http.ListenAndServe(":8080", site)
```