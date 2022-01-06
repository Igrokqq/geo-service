package v1

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vfilipovsky/geo-service/internal/api"
	"github.com/vfilipovsky/geo-service/internal/application"
	"gorm.io/gorm"
)

type TimezoneHandler struct {
	s *application.TimezoneService
}

func newTimezoneHandler(db *gorm.DB) *TimezoneHandler {
	return &TimezoneHandler{
		s: application.NewTimezoneService(db),
	}
}

func AddTimezoneRoutes(r *mux.Router, db *gorm.DB) {
	h := newTimezoneHandler(db)

	r.HandleFunc(api.V1Timezone+"/{id}", h.GetById).Methods(http.MethodGet)
	r.HandleFunc(api.V1Timezone+"/{name}", h.GetByName).Methods(http.MethodGet)
	r.HandleFunc(api.V1Timezone, h.List).Methods(http.MethodGet)
	r.HandleFunc(api.V1Timezone, h.Create).Methods(http.MethodPost)
	r.HandleFunc(api.V1Timezone+"/{id}", h.Update).Methods(http.MethodPut)
	r.HandleFunc(api.V1Timezone+"/{id}", h.Delete).Methods(http.MethodDelete)
}

func (h *TimezoneHandler) List(w http.ResponseWriter, _ *http.Request) {
	list, err := h.s.List()

	if err != nil {
		api.Respond(w, err, http.StatusInternalServerError)
		return
	}

	api.Respond(w, &list, http.StatusOK)
}

func (h *TimezoneHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var err error

	// todo need request parser and validator with human read messages
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		api.Respond(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.s.Delete(uint(i)); err != nil {
		api.Respond(w, err.Error(), http.StatusOK)
		return
	}

	api.Respond(w, "", http.StatusOK)
}

func (h *TimezoneHandler) Update(w http.ResponseWriter, r *http.Request) {
	//todo: parse and validate request
	//todo: create dto
	api.Respond(w, "", http.StatusNotImplemented)
}

func (h *TimezoneHandler) Create(w http.ResponseWriter, r *http.Request) {
	//todo: parse and validate request
	//todo: create dto
	api.Respond(w, "", http.StatusNotImplemented)
}

func (h *TimezoneHandler) GetById(w http.ResponseWriter, r *http.Request) {
	//todo: parse and validate request
	api.Respond(w, "", http.StatusNotImplemented)
}

func (h *TimezoneHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	//todo: parse and validate request
	api.Respond(w, "", http.StatusNotImplemented)
}
