package main

import "github.com/HackYourCareer/SmartKickers/internal/controller/server/server"

const serverAddress = "0.0.0.0:3000"

func main() {
	s := server.New(serverAddress)
	s.Start()
}
