package main

import (
	"fmt"
	"net/http"
)

func mainpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Golang")

}

func main() {
	http.HandleFunc("/", mainpage)
	fmt.Println("port 8080 connected.")

	http.ListenAndServe(":8080", nil)

}
