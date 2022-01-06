package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vfilipovsky/geo-service/internal/config"
	"github.com/vfilipovsky/geo-service/internal/modules/timezone"
	"github.com/vfilipovsky/geo-service/internal/storage"
	"gorm.io/gorm"
)

type app struct {
	db     *gorm.DB
	config *config.Config
}

func newApp() (app, error) {
	if err := godotenv.Load(".env"); err != nil {
		return app{}, err
	}

	cfg := config.Init()
	db, err := storage.NewConnection(cfg.DbCredentials)

	if err != nil {
		return app{}, err
	}

	if err := storage.Migrate(db); err != nil {
		return app{}, err
	}

	return app{
		db:     db,
		config: cfg,
	}, nil
}

func (a app) start() error {
	r := mux.NewRouter()

	timezone.AddRoutes(r, a.db)

	logrus.Println("Listening on port :" + a.config.Http.Port)
	return http.ListenAndServe(":"+a.config.Http.Port, r)
}

func main() {
	a, err := newApp()

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Fatal(a.start())
}
