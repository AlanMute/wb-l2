package handler

import (
	"dev11/internal/core"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := GetPostEvent(r)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	id := h.s.Event.Create(event)
	jsonData, err := json.Marshal(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := GetPostEvent(r)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = h.s.Event.Update(event)

	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	err = json.NewEncoder(w).Encode("Update successful")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	Idstr := r.FormValue("Id")
	id, err := strconv.Atoi(Idstr)

	if err != nil {
		http.Error(w, fmt.Errorf("Id must be integer!").Error(), 400)
		return
	}

	err = h.s.Event.Delete(uint(id))

	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	err = json.NewEncoder(w).Encode("Delete successful")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (h *Handler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	date := r.FormValue("date")

	validDate, err := time.Parse(time.RFC3339, date+"T00:00:00Z")
	if err != nil {
		http.Error(w, fmt.Errorf("Invalid date").Error(), 400)
		return
	}

	events, err := h.s.Event.GetByDay(validDate)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	jsonData, err := json.Marshal(events)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func (h *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	date := r.FormValue("date")

	validDate, err := time.Parse(time.RFC3339, date+"T00:00:00Z")
	if err != nil {
		http.Error(w, fmt.Errorf("Invalid date").Error(), 400)
		return
	}

	events, err := h.s.Event.GetByWeek(validDate)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	jsonData, err := json.Marshal(events)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	date := r.FormValue("date")

	validDate, err := time.Parse(time.RFC3339, date+"T00:00:00Z")
	if err != nil {
		http.Error(w, fmt.Errorf("Invalid date").Error(), 400)
		return
	}

	events, err := h.s.Event.GetByMonth(validDate)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	jsonData, err := json.Marshal(events)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetPostEvent(r *http.Request) (*core.Event, error) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	date := r.FormValue("date")

	validDate, err := time.Parse(time.RFC3339, date+"T00:00:00Z")
	if err != nil {
		return nil, fmt.Errorf("Invalid date")
	}

	validEvent := &core.Event{
		Name:        name,
		Description: description,
		Date:        validDate,
	}

	return validEvent, nil
}
