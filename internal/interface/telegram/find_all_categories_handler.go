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

func (h *FindAllCategoriesHandler) Handle(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	input := usecase.FindAllCategoriesInput{
		TelegramID: uint32(update.Message.From.ID),
	}

	categories, total, err := h.uc.FindAllCategories(input)
	if err != nil {
		log.Printf("Failed to find categories: %v", err)
		h.bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Ошибка при получении категорий. %v", err)))
		return
	} else {
		log.Printf("Found %d categories", total)
	}

	if total == 0 {
		h.bot.Send(tgbotapi.NewMessage(chatID, "У вас ещё нет категорий!"))
		return
	}

	var buttons [][]tgbotapi.InlineKeyboardButton
	for _, category := range categories {
		// Создаем кнопку для каждой категории
		btn := tgbotapi.NewInlineKeyboardButtonData(category.GetName(), fmt.Sprintf("category:%d", category.GetID()))
		buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(btn))
	}

	// Создаем inline-клавиатуру
	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	// Отправляем сообщение с inline-кнопками
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваши категории:")
	msg.ReplyMarkup = keyboard
	h.bot.Send(msg)

}
