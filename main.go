package main

import (
	"flag"
	"log"
	"net/http"
	"remote/internal/echo"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", "0.0.0.0:3000", "http service addr")

func main() {
	flag.Parse()
	log.SetFlags(0)

	r := mux.NewRouter()
	r.HandleFunc("/", echo.Echo)
	r.HandleFunc("/shot", echo.ReceiveShotMsg)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
