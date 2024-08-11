package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pecet3/secret-msg/controllers"
	"github.com/pecet3/secret-msg/messages"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("Loaded .env")
}

func main() {
	ms := messages.NewMsgServices()
	mux := http.NewServeMux()

	controllers.Run(mux, ms)

	address := "127.0.0.1:8010"
	log.Printf("Starting a server [%s]", address)
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	log.Fatalln(server.ListenAndServe())

}
