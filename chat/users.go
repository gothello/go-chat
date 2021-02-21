package chat 

import (
	"log"
	"encoding/json"
	"github.com/gorilla/websocket"
)

type User struct {
	Username string
	Conn *websocket.Conn
	Global *Chat
}

func (u *User) Read() {
	for {

		_, message, err := u.Conn.ReadMessage()
		if err != nil {
			log.Printf("Error on read message %s\n", err)
			return 
		}

		u.Global.messages <- NewMessage(string(message), u.Username)
	}
}

func (u *User) Write(message *Message) {
	json, _ := json.Marshal(message)

	if err := u.Conn.WriteMessage(websocket.TextMessage, json); err != nil {
		log.Printf("Error on write message: %s", err)
		return
	}
}