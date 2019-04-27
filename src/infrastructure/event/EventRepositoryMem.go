package event

import (
	"errors"
	"app/src/models"
)

type EventRepositoryMem struct {
	events []*models.Event
}

func NewEventRepositoryMem() EventRepository {
	return &EventRepositoryMem{[]*models.Event{}}
}

// store event to repository
func (self *EventRepositoryMem) Store(event *models.Event) EventRepository {
	self.events = append(self.events, event)
	return self
}

func (self EventRepositoryMem) FindById(userId string) (*models.Event, error) {
	for _, val := range self.events {
		if val.EventId == userId {
			return val, nil
		}
	}

	return nil, errors.New("user not found")
}

func (self EventRepositoryMem) EventList() []*models.Event {
	return self.events
}