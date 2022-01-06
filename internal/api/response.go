package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Respond(w http.ResponseWriter, value interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	if statusCode >= http.StatusInternalServerError {
		logrus.Error(value)
		send(w, errorResponse{Message: "internal service error", Code: statusCode})
		return
	}

	send(w, &value)
}

func send(w http.ResponseWriter, value interface{}) {
	if err := json.NewEncoder(w).Encode(value); err != nil {
		logrus.Errorf("failed to send response: %s", err.Error())
	}
}
