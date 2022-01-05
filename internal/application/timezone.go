package application

import (
	"github.com/vfilipovsky/geo-service/internal/domain/timezone"
	"github.com/vfilipovsky/geo-service/internal/infrastructure/persistance"
	"gorm.io/gorm"
)

type TimezoneService struct {
	r timezone.Repository
}

func (th *TimezoneService) Get(id uint) (*timezone.Timezone, error) {
	tz, err := th.r.Find(id)

	if err != nil {
		return nil, err
	}

	return tz, nil
}

func NewTimezoneService(db *gorm.DB) *TimezoneService {
	return &TimezoneService{
		r: persistance.NewTimezoneRepository(db),
	}
}
