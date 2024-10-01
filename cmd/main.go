package main

import (
	"NoteKeeperBot/config"
	"NoteKeeperBot/internal/bot"
	db "NoteKeeperBot/internal/database"
	"NoteKeeperBot/internal/storage"

	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	_ = cfg

	dbConn, err := db.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer dbConn.Close()

	store := storage.NewStorage(dbConn)

	botInstance, err := bot.NewBot(cfg, store)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	botInstance.Start()
}
