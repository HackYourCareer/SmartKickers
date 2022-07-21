package echo

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"remote/pkg/messages"
	"strconv"
)

var gameScore messages.GameScore

//	Create a initial response and send back game id json
func handleInitial(mt int, c *websocket.Conn, dm messages.DispatcherMsg) {
	//	package id in json
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
	switch team {
	case 1:
		gameScore.WhiteScore++
	case 2:
		gameScore.BlueScore++
	}
	log.Println("Team 1 score: " + strconv.Itoa(gameScore.WhiteScore) + " Team 2 score: " + strconv.Itoa(gameScore.BlueScore))
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

func ClientServerConn(w http.ResponseWriter, r *http.Request) {
	c := connect(w, r)
	defer c.Close()
	for {
		messageType, p, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))

		// send to client
		if err := c.WriteMessage(messageType, []byte("Hello from server")); err != nil {
			log.Println(err)
			return
		}
		if err := c.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func reader(c *websocket.Conn) {
	for {
		_, p, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		reset := false
		err = json.Unmarshal(p, &reset)
		if reset {
			gameScore.ResetScore()
		}
	}
}

func SendScoreHandler(w http.ResponseWriter, r *http.Request) {
	c := connect(w, r)
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(c) //	Close connection when infinite loop below exits
	mt := 1
	var previousScore messages.GameScore
	go reader(c)

	for {
		if previousScore.BlueScore != gameScore.BlueScore || previousScore.WhiteScore != gameScore.WhiteScore {
			previousScore = gameScore
			gameScoreMsg, _ := json.Marshal(gameScore)
			err := c.WriteMessage(mt, gameScoreMsg)
			if err != nil {
				panic(err)
			}
		}
	}
}
