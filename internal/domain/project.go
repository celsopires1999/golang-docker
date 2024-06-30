package domain

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Project struct {
	ProjectID   string    `validate:"required"`
	Title       string    `validate:"required"`
	Description string    `validate:"-"`
	StarDate    time.Time `validate:"required"`
	ManagerID   string    `validate:"required"`
	EstimatorID string    `validate:"required"`
	CreatedAt   time.Time `validate:"-"`
	UpdatedAt   time.Time `validate:"-"`
}

type RestoreProjectProps Project

var ErrProjectDomainValidation = errors.New("project domain validation failed")

func NewProject(title string, description string, startDate time.Time, managerID string, estimatorID string) *Project {
	return &Project{
		ProjectID:   uuid.New().String(),
		Title:       title,
		Description: description,
		StarDate:    startDate,
		ManagerID:   managerID,
		EstimatorID: estimatorID,
	}
}

func RestoreProject(props RestoreProjectProps) *Project {
	return &Project{
		ProjectID:   props.ProjectID,
		Title:       props.Title,
		Description: props.Description,
		StarDate:    props.StarDate,
		ManagerID:   props.ManagerID,
		EstimatorID: props.EstimatorID,
		CreatedAt:   props.CreatedAt,
		UpdatedAt:   props.UpdatedAt,
	}

}

func (p *Project) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		return ErrProjectDomainValidation
	}
	return nil
}

func (p *Project) AddMonths(months int) error {
	p.StarDate = p.StarDate.AddDate(0, months, 0)
	return nil
}
