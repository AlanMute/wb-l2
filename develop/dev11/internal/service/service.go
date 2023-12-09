package service

import (
	"dev11/internal/core"
	"sync"
)

type Event interface {
}

type Service struct {
	Event
}

func NewService() *Service {
	return &Service{
		Event: NewEventService,
	}
}

type EventService struct {
	mapEvents map[uint]*core.Event
	last      uint
	mu        *sync.RWMutex
}

func NewEventService() *EventService {
	return &EventService{
		mapEvents: make(map[uint]*core.Event),
		last:      0,
	}
}
