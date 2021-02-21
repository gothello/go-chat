package main


import (

	"flag"
	"github.com/gothello/go-chat/chat"
)

var (
	portRunning = flag.Int("port", 3000, "set port running")
)

func init() {
	flag.Parse()
}

func main() {
	chat.Start(*portRunning)
}
