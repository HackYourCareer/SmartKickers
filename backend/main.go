package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/HackYourCareer/SmartKickers/internal/controller/server"
)

const serverAddress = "0.0.0.0:3000"

func main() {
	log.SetLevel(log.DebugLevel)

	s := server.New(serverAddress)

	err := s.Start()
	if err != nil {
		log.Println(err)
	}
}
