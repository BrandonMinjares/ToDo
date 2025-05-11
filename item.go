package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func (e Event) EditEvent(id int, name *string, description *string, completed *bool) error {
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

func DeleteEvent(id int) error {
	_, exists := todoMap[id]
	if !exists {
		return fmt.Errorf("event with ID %d not found", id)
	}
	delete(todoMap, id)
	return nil
}

func EncodeEvents() error {
	jsonBytes, err := json.MarshalIndent(todoMap, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return err
	}

	// Print to terminal
	fmt.Println(string(jsonBytes))

	// Open file for writing
	file, err := os.OpenFile("big_encode.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write JSON bytes directly
	_, err = file.Write(jsonBytes)
	return err
}
