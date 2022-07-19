package main

import (
	"flag"
	"log"
	"net/http"
	"remote/internal/echo"

	"github.com/gorilla/mux"
)

var nodeAddr = flag.String("addr1", "0.0.0.0:3000", "nodeAddr")
var reactAddr = flag.String("addr2", "0.0.0.0:3006", "reactAddr")

func main() {
	flag.Parse()
	log.SetFlags(0)

	r1 := mux.NewRouter()
	r1.HandleFunc("/", echo.Echo)
	r1.HandleFunc("/shot", echo.ReceiveShotMsg)
	go http.ListenAndServe(*nodeAddr, r1)

	r2 := mux.NewRouter()
	r2.HandleFunc("/csc", echo.ClientServerConn)
	http.ListenAndServe(*reactAddr, r2)
}
