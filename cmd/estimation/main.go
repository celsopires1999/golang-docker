package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/celsopires1999/estimation/configs"
	"github.com/celsopires1999/estimation/internal/infra/db"
	httpHandler "github.com/celsopires1999/estimation/internal/infra/http"
	"github.com/celsopires1999/estimation/internal/infra/repository"
	"github.com/celsopires1999/estimation/internal/usecase"
	"github.com/jackc/pgx/v5"
)

func main() {

	ctx := context.Background()

	configs := configs.LoadConfig(".")
	conn, err := pgx.Connect(ctx, configs.DBConn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	if err := conn.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	txm := db.NewTransactionManager(ctx, conn)
	txm.Register("CostRepo", func(q *db.Queries) any {
		return repository.NewCostRepositoryPostgres(q)
	})

	costUsecase := usecase.NewCreateCostUseCase(txm)

	costsHandler := httpHandler.NewCostsHandler(costUsecase)

	r := http.NewServeMux()
	r.HandleFunc("POST /api/v1/costs", costsHandler.CreateCost)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", configs.Port),
		Handler: r,
	}

	// Channel to listen for operating system signals
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		// Received interrupt signal, starting graceful shutdown
		log.Println("Received interrupt signal, starting graceful shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Error during graceful shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	// Starting the HTTP server
	log.Println("HTTP server running on port", configs.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Error starting HTTP server: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("HTTP server finished")
}
