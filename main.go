package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "0.0.0.0:3000", "http service addr")

var upgrader = websocket.Upgrader{}

var goalsOne = 0
var goalsTwo = 0

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			if strings.Contains(string(err.Error()), "websocket") {
				fmt.Println("Team 1 score: " + strconv.Itoa(goalsOne) + " Team 2 score: " + strconv.Itoa(goalsTwo))
			}
			break
		}
		log.Printf("recv: %s", message)
		if strings.Contains(string(message), "INITIAL") {
			_ = c.WriteMessage(mt, json.RawMessage("{\"start\": \"1\" }"))
		}

		if strings.Contains(string(message), "goal") {
			if strings.Contains(string(message), "1") {
				goalsOne++
			}

			if strings.Contains(string(message), "2") {
				goalsTwo++
			}

		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
