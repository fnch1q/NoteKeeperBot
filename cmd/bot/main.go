package main

import (
	"NoteKeeperBot/config"
	db "NoteKeeperBot/internal/infrastructure/database"
	"NoteKeeperBot/internal/interface/telegram"

	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	_ = cfg

	db, err := db.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	_ = db

	bot, err := telegram.NewBot(cfg, db)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	bot.Start()
}
