package main

import (
	"testing"
	"time"
)

func Test_CreateItem(t *testing.T) {
	want := len(todoList) + 1

	e := Event{
		ID:          1,
		Name:        "Write Go code",
		Description: "Finish the TODO app",
		StartTime:   time.Now(), // Current time here
		Completed:   false,
	}
	got := CreateItem(e)

	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}
