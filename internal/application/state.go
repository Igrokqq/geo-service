package application

import (
	"github.com/vfilipovsky/geo-service/internal/domain/state"
	"gorm.io/gorm"
)

type StateService struct {
	r state.Repository
}

func NewStateService(db *gorm.DB) *StateService {
	return &StateService{
		r: nil,
	}
}
