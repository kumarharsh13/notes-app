package main // Every Go program starts with package main

// Import block — group all imports together
// Standard library imports come first, then external packages separated by a blank line
import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// These are YOUR packages — import path = module_name/folder_name
	"github.com/kumarharsh13/notes-app/note"
	"github.com/kumarharsh13/notes-app/todo"
)

// Interface — defines a contract (what methods a type must have)
// Any type with a Save() error method automatically satisfies this
// No "implements" keyword needed — it's implicit
type saver interface {
	Save() error
}

// Embedded interface — combines saver + Display()
// Any type satisfying outputable must have BOTH Save() and Display()
// Both Note and Todo satisfy this because they have both methods
type outputable interface {
	saver     // pulls in Save() error from saver
	Display() // adds Display() on top
}

// main() is the entry point — Go runs this function first
// It can't return values — use fmt.Println + return for errors
func main() {
	// Functions can return multiple values — a Go specialty
	// This returns two strings (title, content)
	title, content := getNoteData()
	text := getUserInput("Todo Data: ")

	// note.New is a constructor — Capitalized because it's exported from the note package
	// Returns (Note, error) — always check the error immediately
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return // exits main since main() can't return an error
	}

	// := reuses err here because todoData is new
	// Rule: := needs at least one NEW variable on the left side
	todoData, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	// outputData accepts any type that satisfies the outputable interface
	// todoData (Todo) works here because Todo has both Display() and Save()
	err = outputData(todoData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// userNote (Note) also works — same interface, different type
	// This is POLYMORPHISM in Go — one function, multiple types
	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Returns two values — Go functions can return multiple values
// No parentheses needed around return types when using named returns
func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

// Regular function — not attached to any struct (no receiver)
func getUserInput(prompt string) string {
	fmt.Print(prompt)

	// bufio.NewReader wraps os.Stdin for buffered reading
	// fmt.Scan stops at spaces — this reads the full line including spaces
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // '\n' is a rune (single quotes), not a string

	if err != nil {
		return ""
	}

	// Clean up the newline characters from the input
	// Windows uses \r\n, Unix uses \n — handle both
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// Accepts the INTERFACE type, not a concrete type
// This function doesn't know or care if it's a Note or Todo
// It only knows the type has Display() and Save()
func outputData(data outputable) error {
	data.Display()        // calls Display() on whatever type was passed
	return saveData(data) // data also satisfies saver (outputable embeds it)
}

// Accepts the narrower saver interface — only needs Save()
// outputable is passed here because it embeds saver
// A type satisfying a wider interface always satisfies narrower ones too
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("[Error] File saving failed")
		return err // pass the error UP to the caller
	}

	fmt.Println("[Success] File saving successfully..!!!!")
	return nil // nil means no error — everything went fine
}
