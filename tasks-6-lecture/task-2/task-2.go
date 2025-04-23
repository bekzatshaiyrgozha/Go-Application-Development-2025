package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"massage"`
}

func mainpage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello, Golang API!"}
	json.NewEncoder(w).Encode(response)
}
func main() {

	http.HandleFunc("/hello", mainpage)
	fmt.Println("8080 is connected")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
