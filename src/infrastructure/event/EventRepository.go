package event

import (
	"app/src/models"
)

type EventRepository interface {
	Store(event *models.Event) EventRepository
	FindById(eventId string) (*models.Event, error)
	EventList() []*models.Event
}