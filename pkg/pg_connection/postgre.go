package db_connection

import (
	"context"
	"fmt"
	"log"
	"spy-cats/pkg/utils"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type StorageConfig struct {
	Host     string `env:"POSTGRESQL_HOST"`
	Port     string
	Username string `env:"POSTGRESQL_USER"`
	Password string `env:"POSTGRESQL_PASSWORD"`
	Database string `env:"POSTGRESQL_DATABASE"`
	SSLMode  string
}

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttempts int, cfg StorageConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	err = utils.ConnectTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, maxAttempts, 5*time.Second)
	if err != nil {
		log.Fatal("error while trying to connect to postgresql")
	}
	return pool, nil
}
