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

func CreateItem(event Event) int {
	todoList = append(todoList, event)
	return len(todoList)
}
