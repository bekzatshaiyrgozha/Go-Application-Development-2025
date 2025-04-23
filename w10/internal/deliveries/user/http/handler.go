package http

import (
	"context"
	"net/http"
	"time"

	"w10/internal/app/config"
	userSrv "w10/internal/services/user"
	"w10/pkg/reqresp"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	srv     userSrv.Service
	timeout time.Duration
}

func NewHandler(cfg config.HTTPServerConfig, srv userSrv.Service) *Handler {
	return &Handler{
		srv:     srv,
		timeout: time.Duration(cfg.RequestTimeout) * time.Second,
	}
}

func (h *Handler) GetUser(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	id := c.Param("id")

	resp, err := h.srv.GetUserByID(ctx, reqresp.GetUserByIDRequest{ID: id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) context(c echo.Context) (context.Context, context.CancelFunc) {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "request_id", c.Response().Header().Get(echo.HeaderXRequestID))

	return context.WithTimeout(ctx, h.timeout)
}
