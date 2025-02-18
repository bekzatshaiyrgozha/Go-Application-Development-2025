cd cmd/_4

go test -bench . unpack_test.go 

go test -bench . -benchmem unpack_test.go

go test -bench . -benchmem json/*.go

go test -bench . -benchmem  string_test.go

go test -bench . -benchmem  prealloc_test.go 

go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1 unpack_test.go

go tool pprof main.test cpu.out
    top
    list Unpack
    web

go tool pprof main.test mem.out
    top
    alloc_space
    top
    list Unpack
    web
    alloc_objects
    list Unpack
    web

go test -bench . -benchmem  pool_test.go 

