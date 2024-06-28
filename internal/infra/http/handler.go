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
	projectRepo := repository.NewProjectRepositoryPostgres(db.New(conn))

	// UseCases
	costUsecase := usecase.NewCreateCostUseCase(txm)
	createProjectUseCase := usecase.NewCreateProjectUseCase(projectRepo)

	// Services
	personService := service.NewUserService(conn)

	// Handlers
	costsHandler := NewCostsHandler(costUsecase)
	usersHandler := NewPersonsHandler(personService)
	projectHandler := NewProjectsHandler(createProjectUseCase)

	// Routes
	r := http.NewServeMux()
	r.HandleFunc("POST /users", usersHandler.CreatePerson)
	r.HandleFunc("PATCH /users/{personID}", usersHandler.UpdatePerson)
	r.HandleFunc("GET /users/{personID}", usersHandler.GetPerson)
	r.HandleFunc("DELETE /users/{personID}", usersHandler.DeletePerson)
	r.HandleFunc("POST /projects", projectHandler.CreateProject)
	r.HandleFunc("POST /costs", costsHandler.CreateCost)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", r))
	return v1
}
