package main

import (
	"bufio"
	"fmt"
	"os"
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
				ID:          1,
				Name:        name,
				Description: description,
				StartTime:   time.Now(),
				Completed:   false,
			}

			CreateEvent(e)

		case "show":
			events := GetEvents()
			for _, event := range events {
				fmt.Println(event.Name)
			}
		default:
			fmt.Println("Goodbye!")
		}
	}
}
