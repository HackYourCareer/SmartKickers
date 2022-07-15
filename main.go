package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type DispatcherMessage struct {
	MsgType string `json:"type"`
	Origin string `json:"origin"`
	TableId int64 `json:"id"`
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Goal int64 `json:"goal"`
	GameStart int64 `json:"start"`
	GameEnd int64 `json:"end"`
}

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

		var dispatcherMsg DispatcherMessage
		er := json.Unmarshal([]byte(message), &dispatcherMsg)

		if er != nil {
			log.Println(err)
		}

		if dispatcherMsg.MsgType == "INITIAL" {
			message := DispatcherMessage {
				GameStart: time.Now().Unix(),
			}
			msg, err := json.Marshal(message)

			if err != nil {
				log.Fatalln(err)
			}

			c.WriteMessage(mt, msg)
		}
		
		if dispatcherMsg.Goal == 1 {
			goalsOne++
		}

		if dispatcherMsg.Goal == 2 {
			goalsTwo++
		}

		log.Println("Team 1 score: " + strconv.Itoa(goalsOne) + " Team 2 score: " + strconv.Itoa(goalsTwo))
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
