package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	ErrIdMustBeAPositiveNum = errors.New("id must be a positive number")
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Respond(w http.ResponseWriter, value interface{}, statusCode int) {
	send(w, &value, statusCode)
}

func RespondAnError(w http.ResponseWriter, err error, statusCode int) {
	if statusCode >= http.StatusInternalServerError {
		logrus.Error(err)
		send(w, errorResponse{Code: statusCode, Message: "internal service error"}, statusCode)
		return
	}

	send(w, errorResponse{Code: statusCode, Message: err.Error()}, statusCode)
}

func send(w http.ResponseWriter, value interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&value); err != nil {
		logrus.Errorf("failed to send response: %s", err.Error())
	}
}
