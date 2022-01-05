package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vfilipovsky/geo-service/internal/infrastructure/config"
	"github.com/vfilipovsky/geo-service/internal/infrastructure/storage"
)

func main() {
	logrus.Println("Starting...")

	if err := godotenv.Load(".env"); err != nil {
		logrus.Panic(err)
	}

	c := config.Init()
	db, err := storage.NewConnection(c.DC)

	if err != nil {
		logrus.Panic(err)
	}

	if err := storage.Migrate(db); err != nil {
		logrus.Panic(err)
	}
}
