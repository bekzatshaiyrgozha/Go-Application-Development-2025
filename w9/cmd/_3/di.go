package _3

type User struct{}

// Определяем интерфейсы
type UserRepository interface {
	GetByID(id int) (*User, error)
	Save(user *User) error
}

type EmailService interface {
	SendWelcomeEmail(user *User) error
}

// Сервис, который зависит от репозитория и email-сервиса
type UserService struct {
	repo  UserRepository
	email EmailService
}

// Constructor Injection
func NewUserService(repo UserRepository, email EmailService) *UserService {
	return &UserService{
		repo:  repo,
		email: email,
	}
}

func (s *UserService) RegisterUser(user *User) error {
	if err := s.repo.Save(user); err != nil {
		return err
	}

	return s.email.SendWelcomeEmail(user)
}
