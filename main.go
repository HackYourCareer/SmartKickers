package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"remote/pkg/messages"
	"strconv"

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

		var dispatcherMsg messages.DispatcherReq
		er := json.Unmarshal([]byte(message), &dispatcherMsg)

		if er != nil {
			log.Println(err)
		}

		if dispatcherMsg.MsgType == "INITIAL" {
			message := messages.DispatcherRes{
				GameId: dispatcherMsg.TableId,
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
