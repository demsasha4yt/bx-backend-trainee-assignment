package bx

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/service"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store/sqlstore"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)

	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	service := service.New(store)
	srv := newServer(store, service)
	fmt.Printf("%+v\n", config.SMTPConfig)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 100
	config.MaxConnLifetime = time.Second * 15
	config.ConnConfig.LogLevel = pgx.LogLevelTrace
	config.ConnConfig.Logger = logrusadapter.NewLogger(logrus.New())

	return pgxpool.ConnectConfig(context.Background(), config)
}
