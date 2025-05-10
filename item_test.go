package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_CreateEvent(t *testing.T) {
	todoList = []Event{}

	e := Event{
		ID:          1,
		Name:        "Write Go code",
		Description: "Finish the TODO app",
		StartTime:   time.Now(),
		Completed:   false,
	}
	CreateEvent(e)

}

func Test_GetEvents(t *testing.T) {
	for _, v := range todoList {
		fmt.Printf("%s\n", v.Name)
	}
}
