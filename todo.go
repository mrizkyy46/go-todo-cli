package main

import (
	"fmt"
)

type Todo struct {
	ID int
	Task string
	Done bool
}

var todos []Todo

func AddTodo(task string) {
	id := GetLastID() + 1

	todo := Todo{
		ID:   id,
		Task: task,
		Done: false,
	}

	todos = append(todos, todo)
}

func ListTodos() {
	for _, t := range todos {

		status := " "

		if t.Done == true {
			status = "x"
		}

		fmt.Printf("[%s] %d - %s\n", status, t.ID, t.Task)
	}
}

func UpdateTodo(id int, task string) {
	for i := 0; i < len(todos); i++ {
		todo := &todos[i]

		if todo.ID == id {
			todo.Task = task

			return
		}
	}
}

func MarkDone(id int)  {
	for i := 0; i < len(todos); i++ {
		todo := &todos[i]

		if todo.ID == id {
			todo.Done = true

			return
		}
	}
}

func DeleteTodo(id int)  {
	var newTodos []Todo

	for i := 0; i < len(todos); i++ {
		if todos[i].ID != id {
			newTodos = append(newTodos, todos[i])
		}
	}

	todos = newTodos
}

func DeleteAllTodos()  {
	var emptyTodo []Todo
	todos = emptyTodo
}

func GetLastID() int  {
	lastID := 0

	if len(todos) == 0 {
		return lastID
	}

	lastIndex := len(todos) - 1

	return todos[lastIndex].ID
}