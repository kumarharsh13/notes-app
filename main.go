package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kumarharsh13/notes-app/note"
	"github.com/kumarharsh13/notes-app/todo"
)

type saver interface { // Normal Interface
	Save() error
}

type outputable interface { // Embbeded Interface
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	text := getUserInput("Todo Data: ")
	userNote, err := note.New(title, content)
	todoData, err := todo.New(text)

	if err != nil {
		fmt.Println(err)
		return
	}

	todoData.Display()

	err = saveData(todoData)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todoData)
	outputData(userNote)

}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("[Error] File saving failed")
		return err
	}

	fmt.Println("[Success] File saving successfully..!!!!")
	return nil
}
