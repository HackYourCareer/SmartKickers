package server

import (
	"net/http"

	"github.com/HackYourCareer/SmartKickers/internal/controller/server/handlers"
	"github.com/HackYourCareer/SmartKickers/internal/model"

	"github.com/gorilla/mux"
)

type Server interface {
	Start() error
}
type server struct {
	router  *mux.Router
	address string
	game    model.Game
}

func New(addr string) Server {
	serv := server{
		router:  mux.NewRouter(),
		address: addr,
		game:    model.Game{},
	}
	serv.router.HandleFunc("/", handlers.TableMessages(serv.game))
	return serv
}

func (s server) Start() error {
	return http.ListenAndServe(s.address, s.router)
}
