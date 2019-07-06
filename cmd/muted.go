package main

import (
	//	"encoding/json"
	"flag"
	//	"errors"
	//	"html/template"
	//	"fmt"
	//	"io/ioutil"
	"log"
	"net/http"
	//	"regexp"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:4004", "http service addresee")

// todo: need to really check the header:Origin before releasing.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("Upgrade:", err)

		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)

			break
		}

		log.Printf("recv: %s", message)

		err = c.WriteMessage(mt, message)

		if err != nil {
			log.Println("write: ", err)

			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	http.HandleFunc("/", echo)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
