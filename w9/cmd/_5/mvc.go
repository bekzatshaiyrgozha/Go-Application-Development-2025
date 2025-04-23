package _5

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model
type Article struct {
	ID      int
	Title   string
	Content string
	Author  string
}

type ArticleModel struct {
	db *sql.DB
}

func (m *ArticleModel) GetByID(id int) (*Article, error) {
	// Получение статьи из БД
	var article Article
	err := m.db.QueryRow("SELECT id, title, content, author FROM articles WHERE id = ?", id).
		Scan(&article.ID, &article.Title, &article.Content, &article.Author)

	if err != nil {
		return nil, err
	}

	return &article, nil
}

// Controller
type ArticleController struct {
	model *ArticleModel
}

func NewArticleController(db *sql.DB) *ArticleController {
	return &ArticleController{
		model: &ArticleModel{db: db},
	}
}

func (c *ArticleController) GetArticle(w http.ResponseWriter, r *http.Request) {
	// Получение ID из URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	// Получение статьи из модели
	article, err := c.model.GetByID(id)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	// Рендеринг представления
	renderArticle(w, article)
}

// View (упрощенный пример)
func renderArticle(w http.ResponseWriter, article *Article) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}
