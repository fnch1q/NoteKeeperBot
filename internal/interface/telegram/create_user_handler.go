package telegram

import (
	"NoteKeeperBot/internal/usecase"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CreateUserHandler struct {
	uc usecase.CreateUserUseCase
}

func NewCreateUserHandler(uc usecase.CreateUserUseCase) CreateUserHandler {
	return CreateUserHandler{uc: uc}
}

func (h *CreateUserHandler) Handle(update tgbotapi.Update) {
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
