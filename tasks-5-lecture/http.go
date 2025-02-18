package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

// Посттарды сақтау үшін структура
type Post struct {
	Title   string
	Content string
}

var (
	posts []Post       // Барлық посттар тізімі
	mutex sync.RWMutex // Бір уақытта бірнеше адам жаза алмауы үшін
)

// Басты бет ("/")
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)
}

// Блог посттарын шығару ("/posts")
func postsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/posts.html")
	mutex.RLock()
	tmpl.Execute(w, posts)
	mutex.RUnlock()
}

// Жаңа пост қосу ("/new")
func newPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")

		if title != "" && content != "" {
			mutex.Lock()
			posts = append(posts, Post{Title: title, Content: content})
			mutex.Unlock()
		}

		http.Redirect(w, r, "/posts", http.StatusSeeOther)
		return
	}

	tmpl, _ := template.ParseFiles("templates/new.html")
	tmpl.Execute(w, nil)
}

func main() {
	// Маршруттарды қосу
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/new", newPostHandler)

	// Статикалық файлдарды қосу (CSS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("🚀 Сервер 8080 портында жұмыс істеп тұр...")
	http.ListenAndServe(":8080", nil)
}
