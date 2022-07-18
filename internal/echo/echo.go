package echo

import (
	"encoding/json"
	"log"
	"net/http"
	"remote/pkg/messages"
	"strconv"

	"github.com/gorilla/websocket"
)

var goalsWhite, goalsBlue int

//	Create a initial response and send back game id json
func handleInitial(mt int, c *websocket.Conn, dm messages.DispatcherMsg) {
	//	package id in json
	/*id, err := strconv.Atoi(dm.TableId)
	if err != nil {
		log.Fatalln("Initial: ", err)
	}*/
	message := messages.DispatcherResponse{
		GameId: dm.TableId,
	}

	msg, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	//	Send json back through connected websocket
	c.WriteMessage(mt, msg)
}

//	Add point if message had a goal field, and print out score
func handleGoal(team int) {

	if team == 1 {
		goalsWhite++
	}
	if team == 2 {
		goalsBlue++
	}

	log.Println("Team 1 score: " + strconv.Itoa(goalsWhite) + " Team 2 score: " + strconv.Itoa(goalsBlue))
}

//	Upgrade the http to websocket connection and check for errors, return the upgraded connection
func connect(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	var upgrader websocket.Upgrader
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } //	Allow all connections by default
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatalln("upgrade:", err)
	}

	return c
}

func Echo(w http.ResponseWriter, r *http.Request) {
	c := connect(w, r)

	defer c.Close() //	Close connection when infinite loop below exits

	for {
		//	Receive message from connection, check for errors
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		//	Resolve given json into provided struct
		var dm messages.DispatcherMsg
		err = json.Unmarshal([]byte(message), &dm)
		if err != nil {
			log.Println(err)
		}

		//	Establish a connection if initial message is received
		if dm.MsgType == "INITIAL" {
			handleInitial(mt, c, dm)
		}

		if dm.Goal != 0 {
			handleGoal(dm.Goal)
		}
	}

}
