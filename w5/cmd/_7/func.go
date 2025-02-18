package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User2 struct {
	ID     int
	Name   string
	Active bool
}

func IsUserOdd(u *User2) bool {
	return u.ID%2 != 0
}

func main() {
	tmplFuncs := template.FuncMap{
		"OddUser": IsUserOdd,
	}

	tmpl, err := template.
		New("").
		Funcs(tmplFuncs).
		ParseFiles("func.html")
	if err != nil {
		panic(err)
	}

	users := []User2{
		User2{1, "Tlep", true},
		User2{2, "Alim", false},
		User2{3, "Azamat", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "func.html",
			struct {
				Users []User2
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
