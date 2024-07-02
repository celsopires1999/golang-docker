package testutils

import (
	"context"
	"log"
	"path/filepath"
	"runtime"

	"github.com/celsopires1999/estimation/configs"
	"github.com/jackc/pgx/v5"
)

func ConnectDB() (*pgx.Conn, string) {
	ctx := context.Background()
	_, base, _, _ := runtime.Caller(0)
	path := filepath.Dir(filepath.Dir(filepath.Dir(base)))

	configs := configs.LoadConfig(path, ".test")
	conn, err := pgx.Connect(ctx, configs.DBConn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return conn, configs.DBConn
}
