package timezone

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vfilipovsky/geo-service/internal/api"
	"gorm.io/gorm"
)

const V1 = "/api/v1/timezone"

type createRequest struct {
	Name string `schema:"name" validate:"required"`
}

type updateRequest struct {
	Name string `schema:"name" validate:"required"`
}

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
	r.HandleFunc(V1+"/name/{name}", h.GetByName).Methods(http.MethodGet)
	r.HandleFunc(V1, h.List).Methods(http.MethodGet)
	r.HandleFunc(V1, h.Create).Methods(http.MethodPost)
	r.HandleFunc(V1+"/{id}", h.Update).Methods(http.MethodPut)
	r.HandleFunc(V1+"/{id}", h.Delete).Methods(http.MethodDelete)
}

func (h *Handler) List(w http.ResponseWriter, _ *http.Request) {
	api.Respond(w, h.s.List())
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := api.ParseIdToUint(mux.Vars(r)["id"])

	if err != nil {
		api.Respond(w, &api.Response{Error: err, Code: http.StatusBadRequest})
		return
	}

	api.Respond(w, h.s.Delete(id))
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := api.ParseIdToUint(mux.Vars(r)["id"])

	if err != nil {
		api.Respond(w, &api.Response{Error: err, Code: http.StatusBadRequest})
		return
	}

	var updateReq updateRequest

	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		api.Respond(w, &api.Response{Error: err, Code: http.StatusBadRequest})
		return
	}

	api.Respond(w, h.s.Update(id, &updateReq))
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var createReq createRequest

	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		api.Respond(w, &api.Response{Error: err, Code: http.StatusBadRequest})
		return
	}

	if err := api.Validate.Struct(createReq); err != nil {
		api.Respond(w, &api.Response{Error: err, Code: http.StatusBadRequest})
		return
	}

	api.Respond(w, h.s.Create(&createReq))
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := api.ParseIdToUint(mux.Vars(r)["id"])

	if err != nil {
		api.Respond(w, &api.Response{Error: err, Code: http.StatusBadRequest})
		return
	}

	api.Respond(w, h.s.GetById(id))
}

func (h *Handler) GetByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	api.Respond(w, h.s.GetByName(name))
}
