package main

import (
	"flag"
	"log"
	"net/http"
	"remote/internal/echo"

	"github.com/gorilla/mux"
)

var nodeAddr = flag.String("addr", "0.0.0.0:3000", "nodeAddr")
var reactAddr = flag.String("addr2", "0.0.0.0:3006", "reactAddr")

func main() {
	flag.Parse()
	log.SetFlags(0)

	r := mux.NewRouter()
	r.HandleFunc("/", echo.Echo)
	r.HandleFunc("/shot", echo.ReceiveShotMsg)
	go http.ListenAndServe(*nodeAddr, r)

	r2 := mux.NewRouter()
	r2.HandleFunc("/csc", echo.ClientServerConn)
	http.ListenAndServe(*reactAddr, r2)
}
