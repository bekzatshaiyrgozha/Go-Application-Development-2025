package _4

import (
	"errors"
	"time"
)

type UserStatus string

const (
	StatusActive   UserStatus = "active"
	StatusInactive UserStatus = "inactive"
	StatusBlocked  UserStatus = "blocked"
)

// User - бизнес-сущность
type User struct {
	ID        int
	Email     string
	Password  string // хеш пароля
	FirstName string
	LastName  string
	Status    UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Метод для валидации пользователя
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	// Дополнительные проверки...

	return nil
}

// Метод для активации пользователя
func (u *User) Activate() error {
	if u.Status == StatusBlocked {
		return errors.New("cannot activate blocked user")
	}

	u.Status = StatusActive
	u.UpdatedAt = time.Now()

	return nil
}

// Метод для блокировки пользователя
func (u *User) Block() {
	u.Status = StatusBlocked
	u.UpdatedAt = time.Now()
}

// FullName возвращает полное имя пользователя
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
