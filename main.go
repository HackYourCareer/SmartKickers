package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"remote/pkg/messages"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "0.0.0.0:3000", "http service addr")

var upgrader = websocket.Upgrader{}

var goalsOne = 0
var goalsTwo = 0

func echo(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()

	for {
		///test
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		var dispatcherMsg messages.DispatcherMsg
		er := json.Unmarshal([]byte(message), &dispatcherMsg)

		if er != nil {
			log.Println(err)
		}

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

	r := mux.NewRouter()
	r.HandleFunc("/", echo)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
