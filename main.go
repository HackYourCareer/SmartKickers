package main

import (
	"flag"
	"log"
	"net/http"
	"remote/internal"
	"remote/internal/echo"
	"remote/pkg/messages"
	"strings"


	"github.com/gorilla/mux"
)

var addr = flag.String("addr", "0.0.0.0:3000", "http service addr")

var upgrader = websocket.Upgrader{}


func main() {
	flag.Parse()
	log.SetFlags(0)

	r := mux.NewRouter()
	r.HandleFunc("/", echo)
	r.HandleFunc("/shot", internal.ReceiveShotMsg)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
