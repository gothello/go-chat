package chat

import (

	"fmt"
	"log"
	"strings"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/gothello/go-chat/utils"
)

type Chat struct {
	users map[string]*User
	messages chan *Message
	join chan *User
	leave chan *User
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 512,
	WriteBufferSize: 512,
	CheckOrigin: func(r *http.Request) bool {
		fmt.Printf("%v %v%v %v\n", r.Method, r.Host, r.RequestURI, r.Proto)
		return r.Method == http.MethodGet
	},

}

func (c *Chat) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error on websocket connection:", err.Error())
	}

	keys := r.URL.Query()
	username := keys.Get("username")
	if strings.TrimSpace(username) == "" {
		username = fmt.Sprintf("anom-%d", utils.GetRandomInt64())
	}

	user := &User{
		Username: username,
		Conn: conn,
		Global: c,
	}

	c.join <- user

	user.Read()
}

func (c *Chat) Add(user *User) {
	if _, ok := c.users[user.Username]; !ok {
		c.users[user.Username] = user
		log.Printf("Added user: %s, Total: %d\n", user.Username, len(c.users))
	}
}
 
func(c *Chat) Broadcast(message *Message) {
	log.Printf("Broadcast message:%v\n", message)
	for _, user := range c.users {
		user.Write(message)
	}
}

func(c *Chat) Disconnect(user *User) {
	if _, ok := c.users[user.Username]; ok {
		defer user.Conn.Close()
		delete(c.users, user.Username)
		log.Printf("User left the chat: %s, Totla: %d\n", user.Username, len(c.users))
	}
}

func (c *Chat) Run() {

	for {
		select {
		case user := <- c.join:
			c.Add(user)
		case message := <- c.messages:
			c.Broadcast(message)
		case user := <- c.leave:
			c.Disconnect(user)
		}
	}
}

func Start(port int) {

	log.Printf("Service running port %d\n", port)

	c := &Chat{
		users: make(map[string]*User),
		messages: make(chan *Message),
		join: make(chan *User),
		leave: make(chan *User),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to GO-CHAT"))
	})

	http.HandleFunc("/chat", c.Handler)

    go c.Run()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

