package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func jsonPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello Go text"}
	json.NewEncoder(w).Encode(response)
}

func textPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Hello Go text")

}

func main() {
	http.HandleFunc("/json", jsonPage)
	http.HandleFunc("/text", textPage)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error connection")
	}
}
