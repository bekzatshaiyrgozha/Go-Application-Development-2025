cd w5/cmd/_8

go run pprof_1.go
open in browser http://127.0.0.1:8080/
// нагрузочное тестирование через утилиту ab
ab -t 20 -n 1000000000 -c 5 http://127.0.0.1:8080/
// снимаем метрики
./pprof_1.sh

go run pprof_2.go
ab -n 1000 -c 10 http://127.0.0.1:8080/
curl http://localhost:8080/debug/pprof/goroutine?debug=2 -o goroutines.txt
OR open in browser http://localhost:8080/debug/pprof/goroutine?debug=2

go build -o tracing tracing.go && ./tracing
open in browser http://127.0.0.1:8080/
ab -t 300 -n 1000000000 -c 10 http://127.0.0.1:8080/
curl http://localhost:8080/debug/pprof/trace?seconds=8 -o trace.out
go tool trace -http "127.0.0.1:8081" tracing trace.out
open in browser http://127.0.0.1:8081/