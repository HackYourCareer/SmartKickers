package main

import (
	"flag"
	"log"
	"net/http"
	"remote/internal/echo"
	"remote/pkg/messages"
	"strings"


	"github.com/gorilla/mux"
)

var addr = flag.String("addr", "0.0.0.0:3000", "http service addr")

var upgrader = websocket.Upgrader{}


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
