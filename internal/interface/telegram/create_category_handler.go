package telegram

import (
	"NoteKeeperBot/internal/usecase"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CreateCategoryHandler struct {
	uc  usecase.CreateCategoryUseCase
	bot *tgbotapi.BotAPI
}

func NewCreateCategoryHandler(uc usecase.CreateCategoryUseCase, bot *tgbotapi.BotAPI) CreateCategoryHandler {
	return CreateCategoryHandler{
		uc:  uc,
		bot: bot,
	}
}

func (h *CreateCategoryHandler) Handle(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	input := usecase.CreateCategoryInput{
		TelegramID: uint32(update.Message.From.ID),
		Name:       update.Message.Text,
	}

	err := h.uc.CreateCategory(input)
	if err != nil {
		log.Printf("Failed to create category: %v", err)
		h.bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Ошибка при создании категории. %v", err)))
	} else {
		log.Printf("Category %s created successfully", input.Name)
		h.bot.Send(tgbotapi.NewMessage(chatID, "Категория успешно создана!"))
	}
}
