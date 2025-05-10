package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter something (type 'exit' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "create":
			fmt.Print("Enter event name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter event description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			e := Event{
				ID:          len(todoMap) + 1,
				Name:        name,
				Description: description,
				StartTime:   time.Now(),
				Completed:   false,
			}

			CreateEvent(e)

		case "show":
			events := GetEvents()

			if len(events) == 0 {
				fmt.Println("No events")
				return
			}

			for _, event := range events {
				fmt.Println("ID: " + strconv.Itoa(event.ID) + " " + event.Name)
			}

		case "delete":
			fmt.Print("Enter id of event to delete: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)

			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				return
			}

			DeleteEvent(id)

		default:
			fmt.Println("Goodbye!")
		}
	}
}
