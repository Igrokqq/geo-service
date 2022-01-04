package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/vfilipovsky/geo-service/internal/config"
	"github.com/vfilipovsky/geo-service/internal/storage"
)

func main() {
	log.Println("Starting...")

	if err := godotenv.Load(".env"); err != nil {
		log.Panic(err)
	}

	c := config.Init()
	db, err := storage.NewConnection(c.DC)

	if err != nil {
		log.Panic(err)
	}

	if err := db.Migrate(); err != nil {
		log.Panic(err)
	}
}
