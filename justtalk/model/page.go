package model

import (
	"os"
	"path"
	"strconv"
)

type Page struct {
	Title    string
	No       int
	Size     int
	LastNo int
	PrevNo   int
	NextNo   int
	Messages []Message
}

func (page *Page) GetMessages() error {
	startNo := page.No * page.Size
	var newMessages []Message
	for i := 0; i < page.Size; i++ {
		messageNo := startNo + i
		filename := path.Join("data", page.Title, strconv.Itoa(messageNo))
		messageBody, err := os.ReadFile(filename)
		if err != nil {
			continue
		}
		message := Message{ID: messageNo, Body: string(messageBody)}
		newMessages = append(newMessages, message)

	}
	page.Messages = newMessages

	return nil
}
