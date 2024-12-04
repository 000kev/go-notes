package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.kyg/notes/note"
)

func main() {
	
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

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving failed")
		return
	}
	fmt.Println("Saving succeeded")
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