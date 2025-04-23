package connections

import (
	"w10/internal/app/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connections struct {
	DB *sqlx.DB
}

func (c *Connections) Close() {
	if c.DB != nil {
		_ = c.DB.Close()
	}
}

func New(cfg *config.Config) (*Connections, error) {
	//connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	//	cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Database, cfg.DB.SSLMode)
	//
	//db, err := sqlx.Connect("postgres", connStr)
	//if err != nil {
	//	log.Fatalf("failed to connect to db: %v", err)
	//}

	//err = db.Ping()
	//if err != nil {
	//	log.Fatalf("failed to ping db: %v", err)
	//}

	return &Connections{
		DB: nil,
	}, nil
}
