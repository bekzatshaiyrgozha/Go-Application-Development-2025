cd cmd/_3

go run pack/*

go build gen/* && ./codegen pack/unpack.go pack/marshaller.go

go run pack/*