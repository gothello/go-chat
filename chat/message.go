package chat

import "github.com/gothello/go-chat/utils"

type Message struct {
	ID int64 `json:"id"`
	Body string `json:"body"`
	Sender string `json:"sender"`
}

func NewMessage(body string, sender string) *Message {
	return &Message{
		ID: utils.GetRandomInt64(),
		Body: body,
		Sender: sender,
	}
}