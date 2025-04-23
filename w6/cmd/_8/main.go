package main

import (
	"bytes"
	"fmt"
	"net/http"

	"w6/cmd/_8/item"
	"w6/cmd/_8/template"
)

//go:generate hero -source=./template/

var ExampleItems = []*item.Item{
	&item.Item{1, "student", "kbtu@.kz team"},
	&item.Item{2, "username", "freelancer"},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		buffer := new(bytes.Buffer)
		template.Index(ExampleItems, buffer)
		w.Write(buffer.Bytes())
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
