package server

import (
	"net/http"
	"remote/internal/controller/server/handlers"

	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	address string
}

func New(addr string) Server {
	s := Server{}
	s.router = mux.NewRouter()
	s.address = addr
	s.router.HandleFunc("/", handlers.HandleTableMessages)
	//s.router.HandleFunc("/shot", handlers.)
	//s.router.HandleFunc("/csc", handlers.)
	return s
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.address, s.router)
}
