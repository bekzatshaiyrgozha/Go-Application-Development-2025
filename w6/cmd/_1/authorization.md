Анатомия веб-сервиса
* Авторизация/аутентификация

Request: 
``` 
\> GET /secret_url HTTP/1.1  
\> Host: localhost:8081  
\> Accept: */*  
\> Content-Type: application/json  
\> X-Access-Token: 695a19177a2760MzY3
```

Response:  
```
< HTTP/1.1 401 Unauthorized  
< Content-Type: application/json  
{"error": "invalid token"}
```