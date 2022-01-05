package storage

import (
	"fmt"
	"time"

	"github.com/vfilipovsky/geo-service/internal/domain/airport"
	"github.com/vfilipovsky/geo-service/internal/domain/city"
	"github.com/vfilipovsky/geo-service/internal/domain/continent"
	"github.com/vfilipovsky/geo-service/internal/domain/country"
	"github.com/vfilipovsky/geo-service/internal/domain/region"
	"github.com/vfilipovsky/geo-service/internal/domain/state"
	"github.com/vfilipovsky/geo-service/internal/domain/timezone"
	"github.com/vfilipovsky/geo-service/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Migrate - sync db with gorm models
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&continent.Continent{},
		&region.Region{},
		&country.Country{},
		&country.Language{},
		&country.Translation{},
		&country.Currency{},
		&timezone.Timezone{},
		&country.Map{},
		&country.AlternativeSpell{},
		&country.PostalCode{},
		&country.CName{},
		&country.DialId{},
		&country.DialIdSuffix{},
		&country.Position{},
		&country.Flag{},
		&state.State{},
		&city.City{},
		&airport.Airport{},
	); err != nil {
		return err
	}

	return nil
}

// NewConnection - return a new pointer to the Database or an error
func NewConnection(dc config.DatabaseCredentials) (*gorm.DB, error) {
	time.Sleep(time.Second / 2) // wait for database up

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dc.Host, dc.Port, dc.User, dc.Name, dc.Pass, dc.SslMode)

	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}))

	if err != nil {
		return nil, err
	}

	return conn, nil
}
