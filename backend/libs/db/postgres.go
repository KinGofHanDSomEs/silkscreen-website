package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MustPool(user, password, host string, port int, database string) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, database)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		log.Fatalf("ping error: %v", err)
	}

	return pool
}
