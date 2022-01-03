package storage

import (
	"fmt"
	"time"

	"github.com/vfilipovsky/geo-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database - Storage implementation
type Database struct {
	conn *gorm.DB
}

// Migrate - sync db with gorm models
func (d *Database) Migrate() error {
	if err := d.conn.AutoMigrate(); err != nil {
		return err
	}

	return nil
}

// NewConnection - return a new pointer to the Database or an error
func NewConnection(dc config.DatabaseCredentials) (Storage, error) {
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

	return &Database{
		conn: conn,
	}, nil
}
