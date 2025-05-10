package main

import (
	"fmt"
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

func EditEvent(id int, name *string, description *string, completed *bool) error {
	event, exists := todoMap[id]
	if !exists {
		return fmt.Errorf("event with ID %d not found", id)
	}

	if name != nil {
		event.Name = *name
	}
	if description != nil {
		event.Description = *description
	}
	if completed != nil {
		event.Completed = *completed
	}

	todoMap[id] = event
	return nil
}

func DeleteEvent(id int) {
	delete(todoMap, id)
}
