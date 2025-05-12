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

	fmt.Print(`
	Would you like to:
	1. Create a new todo list
	2. Load an existing todo list from a file
	> `)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		// Start fresh
		todoMap = make(map[int]Event)
		nextID = 1
	case "2":
		// Load from file
		path := promptForFilePath()
		err := DecodeEvents(path)
		if err != nil {
			fmt.Println("Failed to load events:", err)
			return
		}
	default:
		fmt.Println("Invalid selection. Exiting.")
		return
	}

	CreateToDO()
}
func CreateToDO() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Type 'create' to create new item, 'show' to show all items, 'edit' to update item, 'delete' to delete an item, 'exit' to quit: ")
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
				continue
			}

			for _, event := range events {
				fmt.Println("ID: " + strconv.Itoa(event.ID) + " " + event.Name + " " + event.Description)
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

				fmt.Println("What would you like to edit? Name, description, or completed")
				field, _ := reader.ReadString('\n')
				field = strings.TrimSpace(field)

				var (
					newName        *string
					newDescription *string
					newCompleted   *bool
				)

				switch field {
				case "name":
					fmt.Print("Enter new name: ")
					val, _ := reader.ReadString('\n')
					val = strings.TrimSpace(val)
					newName = &val
				case "description":
					fmt.Print("Enter new description: ")
					val, _ := reader.ReadString('\n')
					val = strings.TrimSpace(val)
					newDescription = &val
				case "completed":
					val := true
					newCompleted = &val
				default:
					fmt.Println("Unknown field")
					continue
				}

				err = todoMap[id].EditEvent(id, newName, newDescription, newCompleted)
				if err != nil {
					fmt.Println("Error updating event:", err)
				} else {
					fmt.Println("Event updated successfully.")
				}

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

			err = DeleteEvent(id)
			if err != nil {
				fmt.Println("Something went wrong with deletion.")
				return
			}
			fmt.Printf("Event with ID %d deleted.\n", id)

		case "exit":
			fmt.Println("Would you like to save file first? (yes/no)")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)

			if idStr == "yes" {
				EncodeEvents()
				fmt.Println("File created")
			}
			fmt.Println("Goodbye")

		default:
			fmt.Println("Not Valid entry. Please try again.")
			continue
		}
	}
}

func promptForFilePath() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the name of the file: ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)
	return path
}
