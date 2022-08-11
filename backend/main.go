package main

import (
	"github.com/HackYourCareer/SmartKickers/internal/config"
	"github.com/HackYourCareer/SmartKickers/internal/controller/server"
	"github.com/HackYourCareer/SmartKickers/internal/model"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)

	s := server.New(config.ServerAddress, model.NewGame())

	err := s.Start()
	if err != nil {
		log.Error(err)
	}
}
