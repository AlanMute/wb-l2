package service

import (
	"dev11/internal/core"
	"fmt"
	"sync"
	"time"
)

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

func (s *EventService) Create(event *core.Event) uint {
	s.mu.Lock()
	id := s.last
	event.Id = id
	s.mapEvents[id] = event
	s.last++

	s.mu.Unlock()

	return id
}

func (s *EventService) Update(event *core.Event) error {
	s.mu.Lock()

	id := event.Id

	if _, ok := s.mapEvents[id]; ok == false {
		return fmt.Errorf("Event with id %d does not found!", id)
	}

	s.mapEvents[id] = event

	s.mu.Unlock()

	return nil
}

func (s *EventService) Delete(id uint) error {
	s.mu.Lock()

	if _, ok := s.mapEvents[id]; ok == false {
		return fmt.Errorf("Event with id %d does not found!", id)
	}

	delete(s.mapEvents, id)

	s.mu.Unlock()

	return nil
}

func (s *EventService) GetByDay(day time.Time) ([]*core.Event, error) {
	s.mu.RLock()

	var events []*core.Event

	for _, v := range s.mapEvents {
		if v.Date.Year() == day.Year() && v.Date.Month() == day.Month() && v.Date.Day() == day.Day() {
			events = append(events, v)
		}
	}

	s.mu.RUnlock()

	return events, nil
}

func (s *EventService) GetByWeek(day time.Time) ([]*core.Event, error) {
	s.mu.RLock()

	var events []*core.Event

	for _, event := range s.mapEvents {
		if event.Date.Sub(day) >= time.Duration(7*time.Now().Day()) {
			events = append(events, event)
		}
	}

	s.mu.RUnlock()

	return events, nil
}

func (s *EventService) GetByMonth(month time.Time) ([]*core.Event, error) {
	s.mu.RLock()

	events := make([]*core.Event, 0, 2)

	for _, v := range s.mapEvents {
		if v.Date.Year() == month.Year() && v.Date.Month() == month.Month() {
			events = append(events, v)
		}
	}

	s.mu.RUnlock()

	return events, nil
}
