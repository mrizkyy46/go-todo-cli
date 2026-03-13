package main

import (
	"encoding/json"
	"os"
)

var fileName string = "todos.json"

func LoadTodos()  {
	file, err := os.Open(fileName)

	if err != nil {
		return
	}

	defer file.Close()

	var decoder *json.Decoder = json.NewDecoder(file)
	decoder.Decode(&todos)
}

func SaveTodos()  {
	file, err := os.Create(fileName)

	if err != nil {
		return
	}

	defer file.Close()
	
	var encoder *json.Encoder = json.NewEncoder(file)
	encoder.Encode(todos)
}