package server

import (
	"io"
	"net/http"

	"github.com/HackYourCareer/SmartKickers/internal/model"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server interface {
	Start() error
	createResponse(io.Reader) ([]byte, error)
}
type server struct {
	router  *mux.Router
	address string
	game    model.Game
}

func New(addr string, game model.Game) Server {
	serv := server{
		router:  mux.NewRouter(),
		address: addr,
		game:    game,
	}
	serv.router.HandleFunc("/", serv.TableMessagesHandler)
	serv.router.HandleFunc("/score", serv.SendScoreHandler)
	serv.router.HandleFunc("/shot", serv.ShotParametersHandler)
	serv.router.HandleFunc("/reset", serv.ResetScoreHandler).Methods("POST")
	serv.router.HandleFunc("/goal", serv.ManipulateScoreHandler).Methods("POST")
	serv.router.HandleFunc("/stats", serv.ShowStatsHandler).Methods("GET")

	return serv
}

func (s server) Start() error {
	log.WithFields(log.Fields{
		"ip": s.address,
	}).Info("Launching the server.")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	return http.ListenAndServe(s.address, handlers.CORS(corsObj)(s.router))
}
