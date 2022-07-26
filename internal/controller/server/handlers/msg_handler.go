package handlers

import (
	"log"
	"net/http"
)

func HandleMessage(w http.ResponseWriter, r *http.Request) {
	c, err := Connect(w, r)
	if err != nil {
		log.Fatalln(err)
	}

	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(message)
	}
}
