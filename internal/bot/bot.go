package bot

import (
	"NoteKeeperBot/config"
	"NoteKeeperBot/internal/storage"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	API     *tgbotapi.BotAPI
	Storage *storage.Storage
}

func NewBot(cfg *config.Config, store *storage.Storage) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, err
	}

	return &Bot{
		API:     botAPI,
		Storage: store,
	}, nil
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.API.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.API.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			// Обработчики команд
		}
	}
}
