package service

import (
	"dev11/internal/core"
	"sync"
)

type Service struct {
	mapEvents map[uint]*core.Event
	last      uint
	mu        *sync.RWMutex
}
