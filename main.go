package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.kyg/notes/todo"
	"example.kyg/notes/note"
)

// naming convention if your interface only has one method
// checks struct to see if it has a Save method with the right arguments and return typess
type saver interface {
	Save() error
} 
type displayer interface {
	Display()
}

// a combined interface
// type outputtable interface {
// 	Save() error
// 	Display()
// }
// or even better an embedded interface
type outputtable interface {
	saver
	displayer
}
func main() {
	
	// NOTE SECTION

	title, content, err := getNote()
	if err != nil {
		fmt.Print("\n", err, "\n")
		return 
	}

	userNote, err := note.New(title,content)
	if err != nil {
		fmt.Print("\n", err, "\n")
		return 
	}

	err = outputData(userNote)

	if err != nil {
		fmt.Println("Saving note failed")
		return
	}
	fmt.Println("Saving note succeeded")

	// TODO SECTION

	activity, err := getTodo()
	if err != nil {
		fmt.Print("\n", err, "\n")
		return 
	}

	userTodo, err := todo.New(activity)
	if err != nil {
		fmt.Print("\n", err, "\n")
		return 
	}
	err = outputData(userTodo)

	if err != nil {
		fmt.Println("Saving todo failed")
		return
	}
	fmt.Println("Saving todo succeeded")
	
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}

	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func getNote() (string, string, error) {
	title, errTitle := getInput("Note title: ")
	content, errContent := getInput("Note content: ")

	if errTitle != nil {
		fmt.Print(errTitle)
		return "", "", errTitle
	} else if errContent != nil {
		fmt.Print(errContent)
		return "", "", errContent
	}

	return title, content, nil
}

func getTodo() (string, error) {
	content, err := getInput("Add a to-do: ")

	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return content, nil
}

// handling long user input text
func getInput(prompt string) (string, error) {
	
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // single quotes for single character

	if err != nil {
		return "", err
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text, nil
}