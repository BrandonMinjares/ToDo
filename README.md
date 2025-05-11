# ToDo

# Go CLI Todo App

A simple command-line Todo application written in Go. This app allows you to create, view, edit, and delete tasks using an in-memory data store.

## Features

- ✅ Create a new todo item with a name and description  
- 📋 Show all todo items with their ID, name, description, and completion status  
- ✏️ Edit existing items by ID (name, description, or completion status)  
- ❌ Delete items by ID  
- 🕒 Automatically records the creation time of each item  

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or later

### Running the App

Clone the repository and run:

```bash
go run main.go item.go


#### Example Usage

> create
Enter event name: Buy groceries
Enter event description: Eggs, milk, bread

> show
ID: 1
Name: Buy groceries
Description: Eggs, milk, bread
Completed: false

> edit
Enter event you would like to edit by ID: 1
What would you like to edit? Name, description, or completed: completed

> show
ID: 1
Name: Buy groceries
Description: Eggs, milk, bread
Completed: true
