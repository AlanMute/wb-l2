package service

import (
	"dev11/internal/core"
	"time"
)

type Event interface {
	Create(event *core.Event) uint
	Update(event *core.Event) error
	Delete(id uint) error
	GetByDay(day time.Time) ([]*core.Event, error)
	GetByWeek(day time.Time) ([]*core.Event, error)
	GetByMonth(month time.Time) ([]*core.Event, error)
}

type Service struct {
	Event
}

func NewService() *Service {
	return &Service{
		Event: NewEventService(),
	}
}
