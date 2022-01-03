package storage

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/vfilipovsky/geo-service/internal/config"
)

// Database - Storage implementation
type Database struct {
	conn *gorm.DB
}

// Migrate - sync db with gorm models
func (d *Database) Migrate() error {
	if result := d.conn.AutoMigrate(); result.Error != nil {
		return result.Error
	}

	return nil
}

// NewConnection - return a new pointer to the Database or an error
func NewConnection(dc config.DatabaseCredentials) (Storage, error) {
	time.Sleep(time.Second / 2) // wait for database up

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dc.Host, dc.Port, dc.User, dc.Name, dc.Pass, dc.SslMode)

	conn, err := gorm.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	if err := conn.DB().Ping(); err != nil {
		return nil, err
	}

	return &Database{
		conn: conn,
	}, nil
}
