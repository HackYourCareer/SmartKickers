package main

import "remote/internal/controller/server"

const serverAddress = "0.0.0.0:3000"

func main() {
	s := server.New(serverAddress)
	s.Start()
}
