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

	//r.HandleFunc(V1+"/{id}", h.GetById).Methods(http.MethodGet)
	//r.HandleFunc(V1+"/{name}", h.GetByName).Methods(http.MethodGet)
	r.HandleFunc(V1, h.List).Methods(http.MethodGet)
	r.HandleFunc(V1, h.Create).Methods(http.MethodPost)
	//r.HandleFunc(V1+"/{id}", h.Update).Methods(http.MethodPut)
	r.HandleFunc(V1+"/{id}", h.Delete).Methods(http.MethodDelete)
}

func (h *Handler) List(w http.ResponseWriter, _ *http.Request) {
	list, err := h.s.List()

	if err != nil {
		api.RespondAnError(w, err, http.StatusInternalServerError)
		return
	}

	api.Respond(w, &list, http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := api.ParseIdToUint(mux.Vars(r)["id"])

	if err != nil {
		api.RespondAnError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.s.Delete(id); err != nil {
		api.RespondAnError(w, err, http.StatusInternalServerError)
		return
	}

	api.Respond(w, nil, http.StatusOK)
}

//func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
//	//todo: parse and validate request
//	//todo: create dto
//	api.Respond(w, "", http.StatusNotImplemented)
//}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var createReq createRequest

	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		api.RespondAnError(w, err, http.StatusBadRequest)
		return
	}

	if err := api.Validate.Struct(createReq); err != nil {
		api.RespondAnError(w, err, http.StatusBadRequest)
		return
	}

	tz, err := h.s.Create(&createReq)

	if err != nil {
		// figure out how to catch unique index err and respond with message
		// timezone is already exists with this name
		api.RespondAnError(w, err, http.StatusInternalServerError)
		return
	}

	api.Respond(w, tz, http.StatusOK)
}

//func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
//	//todo: parse and validate request
//	api.Respond(w, "", http.StatusNotImplemented)
//}
//
//func (h *Handler) GetByName(w http.ResponseWriter, r *http.Request) {
//	//todo: parse and validate request
//	api.Respond(w, "", http.StatusNotImplemented)
//}
