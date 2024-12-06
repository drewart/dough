package main

import (
	"fmt"
	//"net/http"
	"os"
	//"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	// client := websocket.NewClient
	s := ""
	if len(os.Args) > 0 {
		s = os.Args[1]
	}

	fmt.Printf("Hello World! %s", s)

}
