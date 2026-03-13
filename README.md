# Go Todo CLI

Simple **Todo List CLI application** built using Go.

This project is created to practice fundamental Go concepts such as:

- struct
- slice
- pointer
- file handling
- JSON encoding/decoding
- CLI arguments

Todos are stored locally in a JSON file.

---

# Features

- Add todo
- List todos
- Mark todo as done
- Edit todo
- Delete todo
- Delete all todos
- Persistent storage using JSON file

---

# Project Structure

```
go-todo-cli
│
main.go        # CLI command handler
todo.go        # Todo business logic
storage.go     # File persistence (JSON)
go.mod
todos.json     # Generated automatically
```

---

# Running the Application

Run the program from the project directory:

```
go build
```

You can also pass commands directly:

```
go run . <command>
```

---

# CLI Commands

## Add Todo

Add a new todo.

```
./todo-cli add "belajar go"
```

Example:

```
./todo-cli add "belajar goroutine"
```

---

## List Todos

Show all todos.

```
./todo-cli list
```

Example output:

```
[ ] 1 - belajar go
[x] 2 - belajar gin
```

Legend:

- `[ ]` : not completed
- `[x]` : completed

---

## Mark Todo as Done

Mark a todo as completed.

```
./todo-cli done <id>
```

Example:

```
./todo-cli done 1
```

Output:

```
todo marked done
```

---

## Edit Todo

Update a todo task.

```
./todo-cli edit <id> "new task"
```

Example:

```
./todo-cli edit 1 "belajar golang"
```

Output:

```
todo updated
```

---

## Delete Todo

Delete a specific todo.

```
./todo-cli delete <id>
```

Example:

```
./todo-cli delete 2
```

Output:

```
todo deleted
```

---

## Delete All Todos

Remove all todos.

```
./todo-cli delete-all
```

Output:

```
all todos deleted
```

---

# Data Storage

Todos are saved in:

```
todos.json
```

Example file:

```
[
  {
    "id": 1,
    "task": "belajar go",
    "done": false
  }
]
```

---

# Concepts Practiced

This project demonstrates fundamental Go programming concepts:

- Struct modeling
- Slice manipulation
- Pointer usage
- CLI argument parsing
- JSON serialization
- File persistence

---

# Future Improvements

Possible improvements for this project:

- Add search functionality
- Add task priority
- Add due dates
- Convert CLI into REST API
- Implement repository pattern

---

# License

MIT License
