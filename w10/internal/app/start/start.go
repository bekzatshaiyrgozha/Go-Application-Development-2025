package start

import (
	"context"
	"log"
	"time"

	"w10/internal/app/config"
	"w10/internal/deliveries/user/http"
	userSrv "w10/internal/services/user"
	"w10/pkg/graceful"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HTTP(errs chan<- error, cfg *config.Config, srv userSrv.Service) graceful.Service {
	startType := "http"

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	h := http.NewHandler(cfg.HTTPServer, srv)
	http.Register(e, h)

	go func() {
		errs <- e.Start(cfg.HTTPServer.Addr)
	}()

	return graceful.NewService(startType, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			log.Printf("http shutdown error: %v", err)
		}
	})
}
