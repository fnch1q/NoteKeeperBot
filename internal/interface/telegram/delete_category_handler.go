package telegram

import (
	"NoteKeeperBot/internal/usecase"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DeleteCategoryHandler struct {
	uc  usecase.DeleteCategoryUseCase
	bot *tgbotapi.BotAPI
}

func NewDeleteCategoryHandler(uc usecase.DeleteCategoryUseCase, bot *tgbotapi.BotAPI) DeleteCategoryHandler {
	return DeleteCategoryHandler{
		uc:  uc,
		bot: bot,
	}
}

func (h *DeleteCategoryHandler) Handle(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	input := usecase.DeleteCategoryInput{
		TelegramID: uint32(update.Message.From.ID),
		Name:       update.Message.Text,
	}

	err := h.uc.DeleteCategory(input)
	if err != nil {
		log.Printf("Failed to delete category: %v", err)
		h.bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Ошибка при удалении категории. %v", err)))
	} else {
		log.Printf("Category %s deleted successfully", input.Name)
		h.bot.Send(tgbotapi.NewMessage(chatID, "Категория успешно удалена!"))
	}
}
