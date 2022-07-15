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

		var dispatcherMsg messages.DispatcherReqMsg
		er := json.Unmarshal([]byte(message), &dispatcherMsg)

		if er != nil {
			log.Println(err)
		}

		if dispatcherMsg.MsgType == "INITIAL" {
			message := messages.DispatcherResMsg{
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
