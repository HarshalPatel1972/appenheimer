package apperrors

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func WriteHTTP(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   http.StatusText(status),
		Message: msg,
	})
}
