package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pecet3/secret-msg/messages"
)

const DOMAIN_NAME = "secretmsg.pecet.it"

type msgDto struct {
	Message string `json:"message"`
}

type handlers struct {
	ms *messages.MessagesService
}

func Run(srv *http.ServeMux, ms *messages.MessagesService) {
	h := &handlers{
		ms: ms,
	}

	srv.HandleFunc("/message", h.addMsgController)
}

func (h handlers) addMsgController(w http.ResponseWriter, r *http.Request) {
	var dto msgDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(dto)
	token := h.ms.AddMessage(dto.Message)

	log.Println(token)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(DOMAIN_NAME + "/" + token.String()))
}
