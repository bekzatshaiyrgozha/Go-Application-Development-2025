package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User1 struct {
	ID     int
	Name   string
	Active bool
}

func (u *User1) PrintActive() string {
	if !u.Active {
		return ""
	}
	return "method says user " + u.Name + " active"
}

func main() {
	tmpl, err := template.
		New("").
		ParseFiles("method.html")
	if err != nil {
		panic(err)
	}

	users := []User1{
		User1{1, "Tlep", true},
		User1{2, "Alim", false},
		User1{3, "Azamat", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "method.html",
			struct {
				Users []User1
			}{
				users,
			})
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
