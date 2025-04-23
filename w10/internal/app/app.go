package app

import (
	"log"

	"w10/internal/app/config"
	"w10/internal/app/connections"
	"w10/internal/app/start"
	"w10/internal/app/store"
	userSrv "w10/internal/services/user"
	"w10/pkg/graceful"
)

func Run(filenames ...string) {
	cfg, err := config.New(filenames...)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	conns, err := connections.New(cfg)
	if err != nil {
		log.Fatalf("failed to open connections: %v", err)
	}

	st := store.NewStore(conns)
	srv := userSrv.NewService(st)

	errs := make(chan error, 50)

	grace := graceful.New(
		start.HTTP(errs, cfg, srv),
	)
	grace.Shutdown(errs, log.Default(), conns)
}
