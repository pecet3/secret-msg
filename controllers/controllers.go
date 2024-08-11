package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pecet3/secret-msg/messages"
	"github.com/pecet3/secret-msg/views"
)

const DOMAIN_NAME = "secretmsg.pecet.it"

type msgDto struct {
	Message string `json:"message"`
}

type controllers struct {
	ms *messages.MessagesService
}

func Run(srv *http.ServeMux, ms *messages.MessagesService) {
	h := &controllers{
		ms: ms,
	}

	srv.HandleFunc("/add-message", h.addMsgController)
	srv.HandleFunc("/", h.mainPageController)
}

func (h controllers) addMsgController(w http.ResponseWriter, r *http.Request) {
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

func (h controllers) mainPageController(w http.ResponseWriter, r *http.Request) {
	views.AddMsgPage().Render(r.Context(), w)

}
