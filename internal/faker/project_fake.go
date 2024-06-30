package faker

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/google/uuid"
)

type ProjectFakeBuilder struct {
	ProjectID   string
	Title       string
	Description string
	StartDate   time.Time
	ManagerID   string
	EstimatorID string
	CreatedAt   time.Time
	updatedAt   time.Time
}

func NewProjectFakeBuilder() *ProjectFakeBuilder {
	return &ProjectFakeBuilder{
		ProjectID:   uuid.New().String(),
		Title:       randomdata.SillyName(),
		Description: randomdata.Paragraph(),
		StartDate:   time.Date(randomdata.Number(2020, 2030), time.Month(randomdata.Number(1, 12)), 1, 0, 0, 0, 0, time.UTC),
		ManagerID:   uuid.New().String(),
		EstimatorID: uuid.New().String(),
		CreatedAt:   time.Now(),
		updatedAt:   time.Now(),
	}
}

func (b *ProjectFakeBuilder) WithProjectID(projectID string) *ProjectFakeBuilder {
	b.ProjectID = projectID
	return b
}

func (b *ProjectFakeBuilder) WithTitle(title string) *ProjectFakeBuilder {
	b.Title = title
	return b
}

func (b *ProjectFakeBuilder) WithDescription(description string) *ProjectFakeBuilder {
	b.Description = description
	return b
}

func (b *ProjectFakeBuilder) WithStarDate(starDate time.Time) *ProjectFakeBuilder {
	b.StartDate = starDate
	return b
}

func (b *ProjectFakeBuilder) WithManagerID(managerID string) *ProjectFakeBuilder {
	b.ManagerID = managerID
	return b
}

func (b *ProjectFakeBuilder) WithEstimatorID(estimatorID string) *ProjectFakeBuilder {
	b.EstimatorID = estimatorID
	return b
}

func (b *ProjectFakeBuilder) WithCreatedAt(createdAt time.Time) *ProjectFakeBuilder {
	b.CreatedAt = createdAt
	return b
}

func (b *ProjectFakeBuilder) WithUpdatedAt(updatedAt time.Time) *ProjectFakeBuilder {
	b.updatedAt = updatedAt
	return b
}

func (b *ProjectFakeBuilder) Build() *domain.Project {
	props := domain.RestoreProjectProps{}

	props.ProjectID = b.ProjectID
	props.Title = b.Title
	props.Description = b.Description
	props.StarDate = b.StartDate
	props.ManagerID = b.ManagerID
	props.EstimatorID = b.EstimatorID
	props.CreatedAt = b.CreatedAt
	props.UpdatedAt = b.updatedAt

	project := domain.RestoreProject(props)
	err := project.Validate()
	if err != nil {
		panic(err)
	}
	return project
}
