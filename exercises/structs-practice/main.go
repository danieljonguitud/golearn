package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/structs-practice/note"
	"example.com/structs-practice/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo,err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

func getUserInput(prompt string) string {
		fmt.Printf("%v ", prompt)

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')

		if err != nil {
				return ""
		}

		text = strings.TrimSuffix(text, "\n")
		text = strings.TrimSuffix(text, "\r")

		return text
}

func getNoteData() (string, string) {
		title := getUserInput("Note title:")

		content := getUserInput("Note content:")

		return title, content
}

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println("String:", value)
	}
}

func outputData(data outputtable) error {
	data.Display()
    return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving note failed")
		return err
	}

	fmt.Println("Saving note succeeded!")
	return nil
}
