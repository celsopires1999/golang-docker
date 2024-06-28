package domain

import "context"

type CostRepository interface {
	CreateCost(ctx context.Context, cost *Cost) error
	GetCost(ctx context.Context, costID string) (*Cost, error)
	UpdateCost(ctx context.Context, cost *Cost) error
}

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *Project) error
	GetProject(ctx context.Context, projectID string) (*Project, error)
	UpdateProject(ctx context.Context, project *Project) error
	DeleteProject(ctx context.Context, projectID string) error
}
