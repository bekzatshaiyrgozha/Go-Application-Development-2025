package _5

import "errors"

// Сервис для работы со статьями
type ArticleService struct {
	repo ArticleRepository
}

func NewArticleService(repo ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) GetArticle(id int) (*Article, error) {
	return s.repo.GetByID(id)
}

func (s *ArticleService) CreateArticle(title, content, author string) (*Article, error) {
	article := &Article{
		Title:   title,
		Content: content,
		Author:  author,
	}

	// Валидация
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	// Сохранение
	err := s.repo.Save(article)
	if err != nil {
		return nil, err
	}

	return article, nil
}
