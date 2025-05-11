package main

import (
	"testing"
)

// Helper to reset global state before each test
func resetState() {
	todoMap = make(map[int]Event)
	nextID = 1
}

func TestCreateEvent(t *testing.T) {
	resetState()

	name := "Test Event"
	description := "This is a test"

	event := CreateEvent(name, description)

	// Check ID
	if event.ID != 1 {
		t.Errorf("Expected ID 1, got %d", event.ID)
	}

	// Check Name and Description
	if event.Name != name {
		t.Errorf("Expected name '%s', got '%s'", name, event.Name)
	}
	if event.Description != description {
		t.Errorf("Expected description '%s', got '%s'", description, event.Description)
	}

	// Check Completed is false
	if event.Completed != false {
		t.Errorf("Expected Completed to be false, got true")
	}

	// Check if it was added to todoMap
	saved, exists := todoMap[event.ID]
	if !exists {
		t.Errorf("Event with ID %d not found in todoMap", event.ID)
	}
	if saved != event {
		t.Errorf("Saved event does not match returned event")
	}
}
