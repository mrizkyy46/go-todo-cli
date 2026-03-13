package main

import (
	"fmt"
	"os"
)

func main() {
	LoadTodos()

	args := os.Args
	command := args[1]

	if len(os.Args) < 2 {
		fmt.Println("command required: add | list")
		return
	}

	switch command {
	case "add":
		task := args[2]

		AddTodo(task)
		SaveTodos()
		fmt.Println("Todo added")

	case "edit":
		var id int

		task := args[3]

		fmt.Sscanf(args[2], "%d", &id)

		UpdateTodo(id, task)
		SaveTodos()

		fmt.Println("todo updated")

	case "delete":
		var id int

		fmt.Sscanf(args[2], "%d", &id)

		DeleteTodo(id)
		SaveTodos()

		fmt.Println("Todo deleted")

	case "delete-all":
		DeleteAllTodos()
		SaveTodos()

		fmt.Println("All todos deleted")

	case "list":
		ListTodos()

	case "done":
		var id int

		fmt.Sscanf(args[2], "%d", &id)

		MarkDone(id)
		SaveTodos()
		
		fmt.Println("Todo marked done")

	default:
		fmt.Println("unknown command")
	}
}
