package domain_test

import (
	"errors"
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestUnitProject(t *testing.T) {
	t.Run("should create a project with valid values", func(t *testing.T) {
		faker := testutils.NewProjectFakeBuilder()

		project := domain.NewProject(
			faker.Title,
			faker.Description,
			faker.StartDate,
			faker.ManagerID,
			faker.EstimatorID,
		)

		assert.Equal(t, faker.Title, project.Title)
		assert.Equal(t, faker.Description, project.Description)
		assert.Equal(t, faker.StartDate, project.StarDate)
		assert.Equal(t, faker.ManagerID, project.ManagerID)
		assert.Equal(t, faker.EstimatorID, project.EstimatorID)
	})

	t.Run("should fail to create a project with invalid values", func(t *testing.T) {
		faker := testutils.NewProjectFakeBuilder()

		project := domain.NewProject(
			"",
			"",
			faker.StartDate,
			"",
			"",
		)

		err := project.Validate()
		assert.Error(t, err)
		assert.True(t, errors.Is(err, domain.ErrProjectDomainValidation))
	})

	t.Run("should add months to planned start", func(t *testing.T) {
		project := testutils.NewProjectFakeBuilder().
			WithStarDate(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)).
			Build()

		err := project.AddMonths(5)
		assert.NoError(t, err)

		expectedStartDate := time.Date(2020, time.June, 1, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, expectedStartDate, project.StarDate)
	})
}
