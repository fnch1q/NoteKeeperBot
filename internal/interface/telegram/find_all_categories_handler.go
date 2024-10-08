package telegram

import (
	"NoteKeeperBot/internal/usecase"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type FindAllCategoriesHandler struct {
	uc  usecase.FindAllCategoriesUsecase
	bot *tgbotapi.BotAPI
}

func NewFindAllCategoriesHandler(uc usecase.FindAllCategoriesUsecase, bot *tgbotapi.BotAPI) FindAllCategoriesHandler {
	return FindAllCategoriesHandler{
		uc:  uc,
		bot: bot,
	}
}

func (h *FindAllCategoriesHandler) Handle(update tgbotapi.Update) ([]string, int64) {
	chatID := update.Message.Chat.ID

	input := usecase.FindAllCategoriesInput{
		TelegramID: uint32(update.Message.From.ID),
	}

	categories, total, err := h.uc.FindAllCategories(input)
	if err != nil {
		log.Printf("Failed to find categories: %v", err)
		h.bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Ошибка при получении категорий. %v", err)))
	} else {
		log.Printf("Found %d categories", total)
		h.bot.Send(tgbotapi.NewMessage(chatID, "Ваши категории:!"))
	}

	categoryNames := make([]string, 0, len(categories))
	for _, category := range categories {
		categoryNames = append(categoryNames, category.GetName())
	}

	return categoryNames, total

}
