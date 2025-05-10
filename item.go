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

var todoMap = make(map[int]Event)

func CreateEvent(event Event) Event {
	todoMap[event.ID] = event

	return event
}

func GetEvents() map[int]Event {
	return todoMap
}

func DeleteEvent(id int) {
	delete(todoMap, id)
}
