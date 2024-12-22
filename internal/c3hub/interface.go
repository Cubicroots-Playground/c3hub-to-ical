package c3hub

import (
	"context"
	"time"
)

type Service interface {
	GetMySchedule(context.Context) ([]Event, error)
}

type Event struct {
	ID        string
	Name      string
	Room      string
	StartTime time.Time
	EndTime   time.Time
}

type Room struct {
	ID       string
	Name     string
	Type     string
	Assembly string
}
