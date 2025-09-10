package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func New(title string, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content must be entered")
	}
	return &Note{
		Title:   title,
		Content: content,
	}, nil
}

func (note Note) Store() bool {
	filePath := "note/data.json"
	content, errors := os.ReadFile(filePath)
	if errors != nil {
		fmt.Println("File does not exist. Create it!")
		fmt.Println("note", note)
		noteArr := []Note{
			{
				Title:   note.Title,
				Content: note.Content,
			},
		}
		jsonString, err := json.Marshal(noteArr)
		if err != nil {
			fmt.Println("error:", err)
			return false
		}
		os.WriteFile(filePath, []byte(jsonString), 0666)
		return true
	}
	var newNote []Note
	err := json.Unmarshal(content, &newNote)
	if err != nil {
		fmt.Println("error:", err)
		return false
	}
	updateNote := append(newNote, note)
	jsonData, err := json.Marshal(updateNote)
	if err != nil {
		fmt.Println("error:", err)
		return false
	}
	os.WriteFile(filePath, []byte(jsonData), 0666)
	return true
}
