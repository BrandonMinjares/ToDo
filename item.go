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
var nextID = 1

func CreateEvent(name string, description string) Event {
	event := Event{
		ID:          nextID,
		Name:        name,
		Description: description,
		StartTime:   time.Now(),
		Completed:   false,
	}

	todoMap[nextID] = event
	nextID++
	return event
}

func GetEvents() map[int]Event {
	return todoMap
}

func DeleteEvent(id int) {
	delete(todoMap, id)
}
