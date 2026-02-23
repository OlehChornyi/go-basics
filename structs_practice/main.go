package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	// displayer
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

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

func printSomething(data interface{}) {
	intValue, ok := data.(int)

	if ok {
		fmt.Println("Integer:", intValue+1)
		return
	}

	floatValue, ok := data.(float64)

	if ok {
		fmt.Println("Float:", floatValue+1.0)
		return
	}
	// switch data.(type) {
	// case int:
	// 	fmt.Println("Integer:", data)
	// case float64:
	// 	fmt.Println("Float:", data)
	// default:
	// 	fmt.Println(data)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getTodoData() string {
	return getUserInput("Todo title: ")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
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

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Generics
func add[T int | float64 | string](a, b T) T {
	return a + b
}
