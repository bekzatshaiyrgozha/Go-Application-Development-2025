package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "adminpassword"
	dbname   = "testdb"
)

var db *sql.DB

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// PostgreSQL-ге қосылу
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	// DB-ге пинг жасау
	err = db.Ping()
	if err != nil {
		log.Fatal("DB ping error:", err)
	}
	fmt.Println("Connected to PostgreSQL!")

	// Маршруттарды орнату
	router := mux.NewRouter()
	router.HandleFunc("/", serveHTML).Methods("GET") // HTML бетті көрсету
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// HTML бетті көрсету (GET /)
func serveHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error loading HTML file", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Барлық қолданушыларды шығару (GET /users)
func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Жаңа қолданушы қосу (POST /users)
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
