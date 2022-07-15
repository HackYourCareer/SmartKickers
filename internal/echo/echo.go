package echo

import (
	"encoding/json"
	"log"
	"net/http"
	"remote/pkg/messages"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var goalsWhite, goalsBlue int

func Echo(w http.ResponseWriter, r *http.Request) {
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
			goalsWhite++
		}

		if dispatcherMsg.Goal == 2 {
			goalsBlue++
		}

		log.Println("Team 1 score: " + strconv.Itoa(goalsWhite) + " Team 2 score: " + strconv.Itoa(goalsBlue))
	}
}
