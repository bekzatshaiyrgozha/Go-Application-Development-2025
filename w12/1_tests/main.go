package main

import (

	// "encoding/json"

	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// основные настройки к базе
	dsn := "root:kbtuonelove@tcp(127.0.0.1:3306)/photolist?charset=utf8&interpolateParams=true"
	db, err := sql.Open("mysql", dsn)

	err = db.Ping() // вот тут будет первое подключение к базе
	if err != nil {
		log.Fatalf("cant connect to db, err: %v\n", err)
	}

	h := &PhotolistHandler{
		St:   NewDbStorage(db),
		Tmpl: NewTemplates(),
	}

	http.HandleFunc("/", h.List)
	http.HandleFunc("/upload", h.Upload)

	staticHandler := http.StripPrefix(
		"/images/",
		http.FileServer(http.Dir("./../images")),
	)
	http.Handle("/images/", staticHandler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
