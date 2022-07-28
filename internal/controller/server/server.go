package server

import (
	"net/http"
	"remote/internal/controller/server/handlers"
	"remote/internal/model"

	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	address string
	game    model.Game
}

func New(addr string) Server {
	s := Server{}
	s.router = mux.NewRouter()
	s.address = addr
	s.router.HandleFunc("/", handlers.TableMessages(s.game))
	return s
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.address, s.router)
}
