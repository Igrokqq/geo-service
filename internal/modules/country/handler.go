package country

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/vfilipovsky/geo-service/internal/api"
	"github.com/vfilipovsky/geo-service/internal/service"
)

const V1 = "/api/v1/country"

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
	//r.HandleFunc(V1+"/name/{name}", h.GetByName).Methods(http.MethodGet)
	//r.HandleFunc(V1, h.List).Methods(http.MethodGet)
	//r.HandleFunc(V1, h.Create).Methods(http.MethodPost)
	//r.HandleFunc(V1+"/{id}", h.Update).Methods(http.MethodPut)
	//r.HandleFunc(V1+"/{id}", h.Delete).Methods(http.MethodDelete)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := api.ParseIdToUint(mux.Vars(r)["id"])

	if err != nil {
		api.Respond(w, &service.Response{Error: err, Code: http.StatusBadRequest})
		return
	}

	api.Respond(w, h.s.GetById(id))
}
