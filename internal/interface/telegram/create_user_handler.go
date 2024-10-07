package telegram

import (
	"NoteKeeperBot/internal/usecase"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CreateUserHandler struct {
	uc  usecase.CreateUserUseCase
	bot *tgbotapi.BotAPI
}

func NewCreateUserHandler(uc usecase.CreateUserUseCase, bot *tgbotapi.BotAPI) CreateUserHandler {
	return CreateUserHandler{
		uc:  uc,
		bot: bot,
	}
}

func (h *CreateUserHandler) Handle(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать! Выберите действие:")
	msg.ReplyMarkup = GetMainMenuKeyboard()
	h.bot.Send(msg)

	input := usecase.CreateUserInput{
		TelegramID: uint32(update.Message.From.ID),
		Name:       update.Message.From.UserName,
	}

	err := h.uc.CreateUser(input)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
	} else {
		log.Printf("User %s created successfully", input.Name)
	}
}
