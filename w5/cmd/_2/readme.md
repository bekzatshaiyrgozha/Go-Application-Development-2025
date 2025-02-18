cd w5/cmd/_2

go run http.go
curl http://127.0.0.1:8080/

go run pages.go
curl http://127.0.0.1:8080/
curl http://127.0.0.1:8080/page
curl http://127.0.0.1:8080/pages

go run servehttp.go
curl http://127.0.0.1:8080/
curl http://127.0.0.1:8080/test/

go run mux.go
curl http://127.0.0.1:8080/
curl http://127.0.0.1:8080/test/123


go run servers.go
curl http://127.0.0.1:8080/test/123
curl http://127.0.0.1:8081/test/123