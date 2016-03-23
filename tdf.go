package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
)

var arg string = ""

func WSServer(ws *websocket.Conn) {
	fmt.Println(arg)
	t, err := tail.TailFile(arg, tail.Config{Follow: true})
	if err != nil {
		fmt.Println(err)
	}
	for {
		for line := range t.Lines {
			l := []byte(line.Text)
			ws.Write(l)
		}
	}
}

func main() {
	// Just the first argument
	arg = os.Args[1]

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/w", websocket.Handler(WSServer))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("HTTP Error: " + err.Error())
	}
}
