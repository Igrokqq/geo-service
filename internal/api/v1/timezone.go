package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vfilipovsky/geo-service/internal/api"
	"github.com/vfilipovsky/geo-service/internal/application"
	"gorm.io/gorm"
)

const TZPath = api.V1 + "timezone"

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

	r.HandleFunc(TZPath+"/{id}", h.GetById).Methods(http.MethodGet)
	r.HandleFunc(TZPath+"/{name}", h.GetByName).Methods(http.MethodGet)
	r.HandleFunc(TZPath, h.List).Methods(http.MethodGet)
	r.HandleFunc(TZPath, h.Create).Methods(http.MethodPost)
	r.HandleFunc(TZPath+"/{id}", h.Update).Methods(http.MethodPut)
	r.HandleFunc(TZPath+"/{id}", h.Delete).Methods(http.MethodDelete)
}

func (h *TimezoneHandler) List(w http.ResponseWriter, r *http.Request) {
	api.Respond(w, "not implemented yet", http.StatusNotImplemented)
}

func (h *TimezoneHandler) Delete(w http.ResponseWriter, r *http.Request) {
	api.Respond(w, "not implemented yet", http.StatusNotImplemented)
}

func (h *TimezoneHandler) Update(w http.ResponseWriter, r *http.Request) {
	api.Respond(w, "not implemented yet", http.StatusNotImplemented)
}

func (h *TimezoneHandler) Create(w http.ResponseWriter, r *http.Request) {
	api.Respond(w, "not implemented yet", http.StatusNotImplemented)
}

func (h *TimezoneHandler) GetById(w http.ResponseWriter, r *http.Request) {
	api.Respond(w, "not implemented yet", http.StatusNotImplemented)
}

func (h *TimezoneHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	api.Respond(w, "not implemented yet", http.StatusNotImplemented)
}
