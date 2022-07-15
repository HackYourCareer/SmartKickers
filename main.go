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

func receiveShotMsg(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		var shotMsg messages.ShotMsg
		_ = json.Unmarshal(message, &shotMsg)
		var shotParams messages.ShotParams
		_ = json.Unmarshal(shotMsg.Params[0], &shotParams)

		log.Printf("Shot with speeed of %v\n", shotParams.Speed)

	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	r := mux.NewRouter()

	r.HandleFunc("/", echo)
	r.HandleFunc("/shot", receiveShotMsg)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
