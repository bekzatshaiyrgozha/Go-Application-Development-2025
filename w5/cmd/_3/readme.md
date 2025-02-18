cd w5/cmd/_3

go run get.go
curl http://127.0.0.1:8080/?param=1&key=2

go run post.go
open in browser http://127.0.0.1:8080

go run cookies.go
open in browser http://127.0.0.1:8080
show console

go run headers.go
curl http://127.0.0.1:8080/
curl http://127.0.0.1:8080/ -H "Accept:123"