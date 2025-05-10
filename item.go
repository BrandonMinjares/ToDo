package main

import (
	"time"
)

type Event struct {
	ID          int
	Name        string
	Description string
	StartTime   time.Time
	Completed   bool
}

var todoList []Event

func CreateEvent(event Event) Event {
	todoList = append(todoList, event)
	return event
}

func GetEvents() []Event {
	return todoList
}
