package timezone

import (
	"github.com/gorilla/mux"
	"github.com/vfilipovsky/geo-service/internal/api"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const V1 = "/api/v1/timezone"

type Handler struct {
	s *Service
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{
		s: NewService(db),
	}
}

func AddRoutes(r *mux.Router, db *gorm.DB) {
	h := newHandler(db)

	r.HandleFunc(V1+"/{id}", h.GetById).Methods(http.MethodGet)
	r.HandleFunc(V1+"/{name}", h.GetByName).Methods(http.MethodGet)
	r.HandleFunc(V1, h.List).Methods(http.MethodGet)
	r.HandleFunc(V1, h.Create).Methods(http.MethodPost)
	r.HandleFunc(V1+"/{id}", h.Update).Methods(http.MethodPut)
	r.HandleFunc(V1+"/{id}", h.Delete).Methods(http.MethodDelete)
}

func (h *Handler) List(w http.ResponseWriter, _ *http.Request) {
	list, err := h.s.List()

	if err != nil {
		api.Respond(w, err, http.StatusInternalServerError)
		return
	}

	api.Respond(w, &list, http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	//todo: parse and validate request
	//todo: create dto
	api.Respond(w, "", http.StatusNotImplemented)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	//todo: parse and validate request
	//todo: create dto
	api.Respond(w, "", http.StatusNotImplemented)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	//todo: parse and validate request
	api.Respond(w, "", http.StatusNotImplemented)
}

func (h *Handler) GetByName(w http.ResponseWriter, r *http.Request) {
	//todo: parse and validate request
	api.Respond(w, "", http.StatusNotImplemented)
}
