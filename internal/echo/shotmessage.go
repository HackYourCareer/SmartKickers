package echo

import (
	"encoding/json"
	"log"
	"net/http"
	"remote/pkg/messages"
)

func ReceiveShotMsg(w http.ResponseWriter, r *http.Request) {

	c := connect(w, r)

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

		log.Printf("Shot with speed of %v\n", shotParams.Speed)

	}
}
