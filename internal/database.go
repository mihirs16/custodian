package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupDBConn() *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), "user=custodian host=127.0.0.1 port=5432 dbname=playground-dev")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err)
	}
	return pool
}
