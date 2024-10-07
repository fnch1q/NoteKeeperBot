package telegram

import (
	"NoteKeeperBot/internal/usecase"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler struct {
	bot                   *tgbotapi.BotAPI
	createUserHandler     CreateUserHandler
	createCategoryHandler CreateCategoryHandler
}

func NewMessageHandler(
	bot *tgbotapi.BotAPI,
	createUserUC usecase.CreateUserUseCase,
	createCategoryUC usecase.CreateCategoryUseCase,
) MessageHandler {
	createUserHandler := NewCreateUserHandler(createUserUC, bot)
	createCategoryHandler := NewCreateCategoryHandler(createCategoryUC, bot)
	return MessageHandler{
		bot:                   bot,
		createUserHandler:     createUserHandler,
		createCategoryHandler: createCategoryHandler,
	}
}

func (mh *MessageHandler) HandleMessage(update tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.IsCommand() {
			mh.handleCommand(update)
		} else {
			mh.handleTextMessage(update)
		}
	}
}

func (mh *MessageHandler) handleCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		// Логика для команды /start
		mh.createUserHandler.Handle(update)
	case "add_category":
		mh.createCategoryHandler.Handle(update)

	default:
		log.Printf("Unknown command: %s", update.Message.Command())
	}
}

func (mh *MessageHandler) handleTextMessage(update tgbotapi.Update) {
	// Обработка текстовых сообщений
}
