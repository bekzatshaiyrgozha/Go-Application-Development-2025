Анатомия веб-сервиса
* Валидация входных параметров
```
> GET /post?id='ddfr
< {"error":"id invalid, want int"}
```

```
> GET /post?id=123
    post := new(Post)
    schema.Parse(&post, r.URL)
< {"post":{"id":123}}
```