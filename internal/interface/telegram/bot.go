package telegram

import (
	"NoteKeeperBot/config"
	"NoteKeeperBot/internal/repo"
	"NoteKeeperBot/internal/usecase"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

type Bot struct {
	api            *tgbotapi.BotAPI
	messageHandler *MessageHandler
}

func NewBot(cfg *config.Config, db *gorm.DB) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, err
	}

	err = SetCommandMenu(botAPI)
	if err != nil {
		return nil, err
	}

	userRepo := repo.NewUserDB(db)
	categoryRepo := repo.NewCategoryDB(db)
	createUserUC := usecase.NewCreateUserUseCase(userRepo)
	createCategoryUC := usecase.NewCreateCategoryUseCase(categoryRepo, userRepo)

	messageHandler := NewMessageHandler(botAPI, createUserUC, createCategoryUC)

	return &Bot{
		api:            botAPI,
		messageHandler: &messageHandler,
	}, nil
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.api.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		b.messageHandler.HandleMessage(update)
	}
}
