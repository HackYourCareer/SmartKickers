package server

import (
	"net/http"
	"remote/internal/controller/server/handlers"
	"remote/internal/model"

	"github.com/gorilla/mux"
)

type server struct {
	router  *mux.Router
	address string
	game    model.Game
}

func New(addr string) server {
	s := server{
		router:  mux.NewRouter(),
		address: addr,
		game:    model.Game{},
	}

	s.router.HandleFunc("/", handlers.TableMessages(s.game))

	return s
}

func (s *server) Start() error {
	return http.ListenAndServe(s.address, s.router)
}
