package todo // Separate package = separate folder

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Simpler struct — just one field
// Still exported (capital T) with a JSON tag
type Todo struct {
	Text string `json:"text"`
}

// Constructor — same pattern as note.New
// Validates input, returns (Todo, error)
func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("[Error] Invalid Input")
	}

	return Todo{
		Text: content,
	}, nil
}

// Save method — satisfies the saver interface
// Same pattern as Note.Save() but with a fixed filename
func (todo Todo) Save() error {
	fileName := "todo.json"

	// Rename to data to avoid shadowing the json import
	data, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

// Display method — satisfies outputable interface
// Note and Todo implement the SAME interface with DIFFERENT behavior
// This is polymorphism — outputData() in main.go handles both identically
func (todo Todo) Display() {
	fmt.Println(todo.Text)
}
