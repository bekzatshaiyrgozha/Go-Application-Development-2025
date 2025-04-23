package _5

import "database/sql"

// Интерфейс репозитория
type ArticleRepository interface {
	GetByID(id int) (*Article, error)
	Save(article *Article) error
	Delete(id int) error
}

// Реализация репозитория для PostgreSQL
type PostgresArticleRepository struct {
	db *sql.DB
}

func (r *PostgresArticleRepository) GetByID(id int) (*Article, error) {
	var article Article
	err := r.db.QueryRow("SELECT id, title, content, author FROM articles WHERE id = $1", id).
		Scan(&article.ID, &article.Title, &article.Content, &article.Author)

	if err != nil {
		return nil, err
	}

	return &article, nil
}
