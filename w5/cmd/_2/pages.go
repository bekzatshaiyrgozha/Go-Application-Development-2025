package main

import (
	"fmt"
	"net/http"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Main page")
	func(){
		
	}
}

func main() {
	http.HandleFunc("/page",func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Single page:", r.URL.String())
		})

	http.HandleFunc("/pages/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Multiple pages:", r.URL.String())
		})

	http.HandleFunc("/", handler1)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
