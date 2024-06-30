package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
)

type CreateProjectUseCase struct {
	repo domain.ProjectRepository
}

func NewCreateProjectUseCase(repo domain.ProjectRepository) *CreateProjectUseCase {
	return &CreateProjectUseCase{
		repo: repo,
	}
}

func (uc *CreateProjectUseCase) Execute(ctx context.Context, input CreateProjectInputDTO) (*CreateProjectOutputDTO, error) {

	startDate, err := time.Parse("02-01-2006", input.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format, expected 'DD-MM-YYYY': %w", err)
	}

	description := ""
	if input.Description != nil {
		description = *input.Description
	}

	project := domain.NewProject(input.Title, description, startDate, input.ManagerID, input.EstimatorID)
	if err := uc.repo.CreateProject(ctx, project); err != nil {
		return nil, err
	}

	createdProject, err := uc.repo.GetProject(ctx, project.ProjectID)
	if err != nil {
		return nil, err
	}

	return &CreateProjectOutputDTO{
		ProjectID:   createdProject.ProjectID,
		Title:       createdProject.Title,
		Description: createdProject.Description,
		StartDate:   createdProject.StarDate,
		ManagerID:   createdProject.ManagerID,
		EstimatorID: createdProject.EstimatorID,
		CreatedAt:   createdProject.CreatedAt,
	}, nil
}

type CreateProjectInputDTO struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description" validate:"-"`
	StartDate   string  `json:"start_date" validate:"required"`
	ManagerID   string  `json:"manager_id" validate:"required"`
	EstimatorID string  `json:"estimator_id" validate:"required"`
}

type CreateProjectOutputDTO struct {
	ProjectID   string    `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	ManagerID   string    `json:"manager_id"`
	EstimatorID string    `json:"estimator_id"`
	CreatedAt   time.Time `json:"created_at"`
}
