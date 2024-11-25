package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func SendJSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	json.NewEncoder(w).Encode(response)
}
