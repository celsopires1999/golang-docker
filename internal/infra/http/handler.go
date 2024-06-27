package http

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"

	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/infra/repository"
	"github.com/celsopires1999/estimation/internal/service"
	"github.com/celsopires1999/estimation/internal/usecase"
)

func Handler(ctx context.Context, conn *pgx.Conn) *http.ServeMux {
	// Transaction Manager
	txm := db.NewTransactionManager(ctx, conn)
	txm.Register("CostRepo", func(q *db.Queries) any {
		return repository.NewCostRepositoryPostgres(q)
	})

	// Repositories

	// UseCases
	costUsecase := usecase.NewCreateCostUseCase(txm)

	// Services
	personService := service.NewUserService(conn)

	// Handlers
	costsHandler := NewCostsHandler(costUsecase)
	personsHandler := NewPersonsHandler(personService)

	// Routes
	r := http.NewServeMux()
	r.HandleFunc("POST /costs", costsHandler.CreateCost)
	r.HandleFunc("POST /persons", personsHandler.CreatePerson)
	r.HandleFunc("PATCH /persons/{personID}", personsHandler.UpdatePerson)
	r.HandleFunc("GET /persons/{personID}", personsHandler.GetPerson)
	r.HandleFunc("DELETE /persons/{personID}", personsHandler.DeletePerson)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", r))
	return v1
}
