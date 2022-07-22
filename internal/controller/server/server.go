package server

import (
	"net/http"
	"remote/internal/controller/server/handlers"

	"github.com/gorilla/mux"
)

type server struct {
	router  *mux.Router
	address string
}

func New(addr string) server {
	s := server{}
	s.router = mux.NewRouter()
	s.address = addr
	s.router.HandleFunc("/", handlers.HandleMessage)
	return s
}

func (s *server) Start() error {
	return http.ListenAndServe(s.address, s.router)
}
