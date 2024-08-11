package main

import (
	"log"
	"net/http"

	"github.com/pecet3/secret-msg/handlers"
	"github.com/pecet3/secret-msg/messages"
)

func main() {
	ms := messages.NewMsgServices()
	mux := http.NewServeMux()

	handlers.Run(mux, ms)

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	address := "127.0.0.1:8010"
	log.Printf("Starting a server [%s]", address)
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}
	log.Fatalln(server.ListenAndServe())
}
