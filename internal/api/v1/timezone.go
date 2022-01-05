package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	th := newTimezoneHandler(db)

	r.HandleFunc("/api/v1/timezone/{id}", th.Get).Methods(http.MethodGet)
}

func (th *TimezoneHandler) Get(w http.ResponseWriter, r *http.Request) {
	// todo request validator https://github.com/go-playground/validator
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		// todo
	}

	result, err := th.s.Get(uint(id))

	if err != nil {
		// todo response writer
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		// todo
	}
}
