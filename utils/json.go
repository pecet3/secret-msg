package utils

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, dto interface{}) error {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(dto)
	if err != nil {
		return err
	}
	return nil
}
