package main

import (
	"github.com/HackYourCareer/SmartKickers/internal/controller/server"
	"github.com/HackYourCareer/SmartKickers/internal/model"
	log "github.com/sirupsen/logrus"
)

const serverAddress = "0.0.0.0:3000"

func main() {
	log.SetLevel(log.InfoLevel)

	s := server.New(serverAddress, model.NewGame())

	err := s.Start()
	if err != nil {
		log.Error(err)
	}
}
