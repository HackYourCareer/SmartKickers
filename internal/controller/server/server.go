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
	s := server{}
	s.router = mux.NewRouter()
	s.address = addr
	s.router.HandleFunc("/", handlers.HandleTableMessages(s.game))
	//s.router.HandleFunc("/shot", handlers.)
	//s.router.HandleFunc("/csc", handlers.)
	return s
}

func (s *server) Start() error {
	return http.ListenAndServe(s.address, s.router)
}
