package todo

import (
	"fmt"
	"encoding/json"
	"errors"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf("Your todo has the following content: %v\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName,json,0644)
 }
func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid todo")
	}
	return Todo{
		Text: content,
	}, nil
}