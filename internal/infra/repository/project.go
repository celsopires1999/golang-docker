package repository

import (
	"context"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProjectRepositoryPostgres struct {
	queries *db.Queries
}

func NewProjectRepositoryPostgres(q *db.Queries) *ProjectRepositoryPostgres {
	return &ProjectRepositoryPostgres{
		queries: q,
	}
}

func (r *ProjectRepositoryPostgres) CreateProject(ctx context.Context, project *domain.Project) error {
	err := r.queries.CreateProject(ctx, db.CreateProjectParams{
		ProjectID:   project.ProjectID,
		Title:       project.Title,
		Description: pgtype.Text{String: project.Description, Valid: true},
		StartDate:   pgtype.Date{Time: project.StarDate, Valid: true},
		ManagerID:   project.ManagerID,
		EstimatorID: project.EstimatorID,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	return err
}

func (r *ProjectRepositoryPostgres) GetProject(ctx context.Context, projectID string) (*domain.Project, error) {
	projectModel, err := r.queries.GetProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	props := domain.RestoreProjectProps{
		ProjectID:   projectModel.ProjectID,
		Title:       projectModel.Title,
		Description: projectModel.Description.String,
		StarDate:    projectModel.StartDate.Time,
		ManagerID:   projectModel.ManagerID,
		EstimatorID: projectModel.EstimatorID,
		CreatedAt:   projectModel.CreatedAt.Time,
		UpdatedAt:   projectModel.UpdatedAt.Time,
	}

	project := domain.RestoreProject(props)
	err = project.Validate()
	if err != nil {
		return nil, err
	}

	return domain.RestoreProject(props), nil
}

func (r *ProjectRepositoryPostgres) UpdateProject(ctx context.Context, project *domain.Project) error {
	return r.queries.UpdateProject(ctx, db.UpdateProjectParams{
		ProjectID:   project.ProjectID,
		Title:       project.Title,
		Description: pgtype.Text{String: project.Description, Valid: true},
		StartDate:   pgtype.Date{Time: project.StarDate, Valid: true},
		ManagerID:   project.ManagerID,
		EstimatorID: project.EstimatorID,
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
}

func (r *ProjectRepositoryPostgres) DeleteProject(ctx context.Context, projectID string) error {
	return r.queries.DeleteProject(ctx, projectID)
}
