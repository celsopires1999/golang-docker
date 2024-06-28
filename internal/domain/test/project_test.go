package domain_test

import (
	"errors"
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestProject(t *testing.T) {
	t.Run("should create a project with valid values", func(t *testing.T) {
		project := domain.NewProject(
			fixtureProject.title,
			fixtureProject.description,
			fixtureProject.startDate,
			fixtureProject.managerID,
			fixtureProject.estimatorID,
		)

		assert.Equal(t, fixtureProject.title, project.Title)
		assert.Equal(t, fixtureProject.description, project.Description)
		assert.Equal(t, fixtureProject.startDate, project.StarDate)
		assert.Equal(t, fixtureProject.managerID, project.ManagerID)
		assert.Equal(t, fixtureProject.estimatorID, project.EstimatorID)
	})

	t.Run("should fail to create a project with invalid values", func(t *testing.T) {

		project := domain.NewProject(
			"",
			"",
			fixtureProject.startDate,
			"",
			"",
		)

		err := project.Validate()
		assert.Error(t, err)
		assert.True(t, errors.Is(err, domain.ErrProjectDomainValidation))
	})

	t.Run("should add months to planned start", func(t *testing.T) {
		project := domain.NewProject(
			fixtureProject.title,
			fixtureProject.description,
			time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			fixtureProject.managerID,
			fixtureProject.estimatorID,
		)

		err := project.AddMonths(5)
		assert.NoError(t, err)

		expectedStartDate := time.Date(2020, time.June, 1, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, expectedStartDate, project.StarDate)
	})
}
