package core

import "time"

type Event struct {
	Id          uint      `json:"Id"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Date        time.Time `json:"Date"`
}
