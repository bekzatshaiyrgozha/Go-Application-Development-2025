cd w5/cmd/_5

go run file_upload.go
open in browser http://127.0.0.1:8080
curl -v -X POST -H "Content-Type: application/json" -d '{"id": 2, "user": "someuser"}' http://localhost:8080/raw_body