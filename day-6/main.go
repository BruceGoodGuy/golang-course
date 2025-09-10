package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title := GetInput("Enter title: ")
	content := GetInput("Enter content: ")
	var noteData *note.Note
	noteData, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error write note:", err)
		return
	}

	noteData.Store()
}

func GetInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		text = ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
