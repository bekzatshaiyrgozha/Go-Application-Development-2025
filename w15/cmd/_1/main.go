package main

import (
	"net/http"

	_ "w15/cmd/_1/docs" // Импорт автогенерированной документации

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Example API
// @version 1.0
// @description This is a sample server using Echo framework.
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api/v1
func main() {
	e := echo.New()

	// Группы маршрутов API
	v1 := e.Group("/api/v1")
	{
		users := v1.Group("/users")
		users.GET("", getUsers)
		users.POST("", createUser)
		users.GET("/:id", getUser)
	}

	// Маршрут для документации Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

// GetUsers godoc
// @Summary Get users
// @Description Get list of users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c echo.Context) error {
	// Реализация
	users := []User{
		{ID: "1", Username: "user1", Email: "user1@example.com"},
	}
	return c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User data"
// @Success 201 {object} User
// @Failure 400 {object} ErrorResponse
// @Router /users [post]
func createUser(c echo.Context) error {
	// Создание нового пользователя
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
	}

	// Здесь была бы логика создания пользователя

	return c.JSON(http.StatusCreated, user)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func getUser(c echo.Context) error {
	id := c.Param("id")
	// Здесь была бы логика получения пользователя по ID

	user := User{
		ID:       id,
		Username: "username",
		Email:    "user@example.com",
	}

	return c.JSON(http.StatusOK, user)
}

// User представляет модель пользователя
// @Description User account information
type User struct {
	ID       string `json:"id" example:"1"`
	Username string `json:"username" example:"john_doe"`
	Email    string `json:"email" example:"john@example.com"`
}

// ErrorResponse представляет стандартный ответ об ошибке
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad request"`
}
