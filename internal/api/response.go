package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Response struct {
	Error error
	Code  int
	Value interface{}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Respond(w http.ResponseWriter, resp *Response) {
	if resp.Error == nil {
		send(w, &resp.Value, resp.Code)
		return
	}

	if resp.Code >= http.StatusInternalServerError {
		logrus.Error(resp.Error)
		send(w, errorResponse{Code: resp.Code, Message: "internal service error"}, resp.Code)
		return
	}

	send(w, errorResponse{Code: resp.Code, Message: resp.Error.Error()}, resp.Code)
}

func send(w http.ResponseWriter, value interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&value); err != nil {
		logrus.Errorf("failed to send response: %s", err.Error())
	}
}
