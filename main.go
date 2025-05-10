package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Type 'create' to create new item, 'show' to show all items, 'delete' to delete an item,'exit' to quit): ")
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

			CreateEvent(name, description)

		case "show":
			events := GetEvents()

			if len(events) == 0 {
				fmt.Println("No events")
				return
			}

			for _, event := range events {
				fmt.Println("ID: " + strconv.Itoa(event.ID) + " " + event.Name)
			}
		case "edit":
			fmt.Print("Enter event you would like to edit by ID: ")
			ID, _ := reader.ReadString('\n')
			ID = strings.TrimSpace(ID)
			id, err := strconv.Atoi(ID)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				return
			}

			if _, exists := todoMap[id]; exists {
				fmt.Println("Id: " + ID)
				fmt.Println("Name: " + todoMap[id].Name)
				fmt.Println("Description: " + todoMap[id].Description)

				if !todoMap[id].Completed {
					fmt.Println("Completed: No")
				} else {
					fmt.Println("Completed: Yes")
				}

				fmt.Println("What would you like to edit?")
			} else {
				fmt.Println("ID not found")
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
