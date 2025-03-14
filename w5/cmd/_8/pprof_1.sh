curl http://127.0.0.1:8080/debug/pprof/heap -o mem_out.txt
curl http://127.0.0.1:8080/debug/pprof/profile?seconds=5 -o cpu_out.txt

go build -o pprof_1 pprof_1.go 
go tool pprof -svg -alloc_objects pprof_1 mem_out.txt > mem_ao.svg
go tool pprof -svg pprof_1 cpu_out.txt > cpu.svg
