package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	v1 "github.com/vfilipovsky/geo-service/internal/api/v1"
	"github.com/vfilipovsky/geo-service/internal/infrastructure/config"
	"github.com/vfilipovsky/geo-service/internal/infrastructure/storage"
	"gorm.io/gorm"
)

type app struct {
	db *gorm.DB
}

func newApp() (app, error) {
	if err := godotenv.Load(".env"); err != nil {
		return app{}, err
	}

	c := config.Init()
	db, err := storage.NewConnection(c.DC)

	if err != nil {
		return app{}, err
	}

	if err := storage.Migrate(db); err != nil {
		return app{}, err
	}

	return app{
		db: db,
	}, nil
}

func (a app) start() error {
	logrus.Println("Starting server")
	logrus.Println("Adding routes")

	r := mux.NewRouter()

	v1.AddTimezoneRoutes(r, a.db)

	logrus.Println("Serving on port :8080") // todo move app configs
	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}

	return nil
}

func main() {
	s, err := newApp()

	if err != nil {
		logrus.Panic(err)
	}

	if err := s.start(); err != nil {
		logrus.Panic(err)
	}
}
