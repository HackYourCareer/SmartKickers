package server

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router  *mux.Router
	address *string
}

func New(addr string) server {
	s := server{}
	s.router = mux.NewRouter()
	s.address = flag.String("address", addr, "Address of the server")
	return s
}

func (s *server) Start() error {
	http.ListenAndServe(*s.address, s.router)
	return nil
}
