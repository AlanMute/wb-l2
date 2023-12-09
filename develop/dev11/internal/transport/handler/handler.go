package handler

import (
	"dev11/internal/service"
	"net/http"
)

type Handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		s: s,
	}
}

// Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", h.MiddlewareLogger(h.CreateEvent))
	mux.HandleFunc("/update_event", h.MiddlewareLogger(h.UpdateEvent))
	mux.HandleFunc("/delete_event", h.MiddlewareLogger(h.DeleteEvent))
	mux.HandleFunc("/events_for_day", h.MiddlewareLogger(h.EventsForDay))
	mux.HandleFunc("/events_for_week", h.MiddlewareLogger(h.EventsForWeek))
	mux.HandleFunc("/events_for_month", h.MiddlewareLogger(h.EventsForMonth))

	return mux
}
