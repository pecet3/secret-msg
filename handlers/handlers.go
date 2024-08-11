package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pecet3/secret-msg/messages"
	"github.com/pecet3/secret-msg/utils"
)

const DOMAIN_NAME = "secretmsg.pecet.it"

type msgDto struct {
	Message     string `json:"message"`
	IsEncrypted bool   `json:"is_encrypted"`
}

type responseDto struct {
	Token string `json:"token"`
}

type handlers struct {
	ms *messages.MessagesService
}

func Run(mux *http.ServeMux, ms *messages.MessagesService) {
	h := &handlers{
		ms: ms,
	}

	mux.HandleFunc("POST /message", h.handleAddMessage)
	mux.HandleFunc("GET /message/{id}", h.handleGetMessage)
}

func (h handlers) handleAddMessage(w http.ResponseWriter, r *http.Request) {
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

func (h handlers) handleGetMessage(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("id")
	if token == "" {
		http.Error(w, "empty token", http.StatusBadRequest)
		return
	}
	msg, err := h.ms.GetMessage(token)
	if err != nil {
		http.Error(w, "no message", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(msg)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
