package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pecet3/secret-msg/messages"
	"github.com/pecet3/secret-msg/utils"
	"github.com/pecet3/secret-msg/views"
)

const DOMAIN_NAME = "secretmsg.pecet.it"

type msgDto struct {
	Message     string `json:"message"`
	IsEncrypted bool   `json:"is_encrypted"`
}

type responseDto struct {
	Token string `json:"token"`
}

type controllers struct {
	ms *messages.MessagesService
}

func Run(mux *http.ServeMux, ms *messages.MessagesService) {
	c := &controllers{
		ms: ms,
	}

	mux.HandleFunc("/read/{id}", c.readPageController)
	mux.HandleFunc("/", c.mainPageController)
}

func (c controllers) mainPageController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		views.MainPage().Render(r.Context(), w)
	}
	if r.Method == http.MethodPost {
		var dto msgDto
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println(dto)
		token := c.ms.AddMessage(dto.Message)
		log.Println(token)

		response := responseDto{
			Token: token,
		}
		err = utils.SendJson(w, response)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
}

func (c controllers) readPageController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		token := r.PathValue("id")
		if token == "" {
			http.Error(w, "empty token", http.StatusBadRequest)
			return
		}
		msg, err := c.ms.GetMessage(token)
		if err != nil {
			http.Error(w, "no message", http.StatusBadRequest)
			return
		}
		err = c.ms.DeleteMessage(token)
		if err != nil {
			return
		}
		views.ReadPage(msg.Content).Render(r.Context(), w)
	}
}
