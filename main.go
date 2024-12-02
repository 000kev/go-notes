package main

import (
	"fmt"
	"errors"
)

func main() {
	
	title, content, err := getNote()
	if err != nil {
		return
	}

	fmt.Println("You entereted the title", title, "\nYou entered the content", content)
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

func getInput(prompt string) (string, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)

	if input == "" {
		return "", errors.New("invalid input")
	}
	return input, nil
}