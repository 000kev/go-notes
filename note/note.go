package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// keeping the fields lowercase keeps them private i.e. only directly accessible within this package
// struct tags are metadata for struct fields so any packages that uses the metadata will look here to use these values
type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the follwing content:\n\n%v\n", note.Title, note.Content)
}

// accessors don't need a pointer to the data
// but mutators do
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note) // only makes json data from publicly available fields
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("invalid input")
	}
	return &Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

