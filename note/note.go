package note // Matches the folder name — all files in note/ must say this

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Struct — groups related data together
// Capitalized field names = exported (visible outside this package)
// Struct tags (json:"title") tell the JSON encoder what keys to use
// Without tags, JSON keys would be "Title", "Content", "CreatedAt"
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Constructor function — Go convention is New or NewTypeName
// Capitalized = exported, so main.go can call note.New()
// Returns (Note, error) — the Go pattern for "this might fail"
func New(title, content string) (Note, error) {
	// When multiple params share a type, you can write (title, content string)
	// instead of (title string, content string)

	if title == "" || content == "" {
		// Return zero-value Note{} on failure + an error
		// Zero value = all fields set to their defaults (empty string, zero time)
		return Note{}, errors.New("[Error] Invalid Input")
	}

	// Struct literal — field: value syntax
	// time.Now() captures the current timestamp
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil // nil = no error
}

// Method with VALUE receiver (note Note)
// Gets a COPY of the Note — can read but can't modify the original
// This satisfies the saver interface's Save() error requirement
func (note Note) Save() error {
	// strings.ReplaceAll — replace every space with underscore
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	// strings.ToLower — lowercase everything, then add .json
	fileName = strings.ToLower(fileName) + ".json"

	// json.Marshal converts struct → JSON bytes
	// Uses the struct tags to decide key names
	// Rename the variable to avoid shadowing the json import
	data, err := json.Marshal(note)

	if err != nil {
		return err // pass the error up
	}

	// os.WriteFile writes bytes to a file
	// 0644 = Unix file permissions (owner read/write, others read)
	// Creates the file if it doesn't exist, overwrites if it does
	return os.WriteFile(fileName, data, 0644)
}

// Another value receiver method
// Satisfies the Display() requirement of the outputable interface
// Together with Save(), Note now satisfies outputable — automatically
func (note Note) Display() {
	// Printf = formatted print
	// %v = default format for any value
	fmt.Printf("Your note titled %v has following content:\n %v\n", note.Title, note.Content)
}
