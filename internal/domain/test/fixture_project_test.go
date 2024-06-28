package domain_test

import (
	"time"

	"github.com/google/uuid"
)

var fixtureProject = struct {
	title       string
	description string
	startDate   time.Time
	managerID   string
	estimatorID string
}{
	title:       "Project Title",
	description: "Project Description",
	startDate:   time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	managerID:   uuid.New().String(),
	estimatorID: uuid.New().String(),
}
