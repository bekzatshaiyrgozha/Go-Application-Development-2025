package http

import "github.com/labstack/echo/v4"

func Register(e *echo.Echo, h *Handler) {
	e.GET("/user/:id", h.GetUser)
}
