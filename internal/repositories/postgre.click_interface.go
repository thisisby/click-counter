package repositories

import (
	"click-counter/internal/models"
	"time"
)

type ClickRepository interface {
	Save(id string) error
	FindByIdInRange(id string, startTimestamp, endTimestamp time.Time) ([]models.Click, error)
}
