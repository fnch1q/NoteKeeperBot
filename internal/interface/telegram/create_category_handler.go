package telegram //Переделать что то не то

import (
	"NoteKeeperBot/internal/usecase"
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CreateCategoryHandler struct {
	uc         usecase.CreateCategoryUseCase
	bot        *tgbotapi.BotAPI
	waitingFor sync.Map // Для управления доступом к карте ожидания
}

func NewCreateCategoryHandler(uc usecase.CreateCategoryUseCase, bot *tgbotapi.BotAPI) CreateCategoryHandler {
	return CreateCategoryHandler{
		uc:  uc,
		bot: bot,
	}
}

func (h *CreateCategoryHandler) Handle(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	waiting, _ := h.waitingFor.Load(chatID)

	if waiting != nil && waiting.(bool) {
		if update.Message.Text == "" {
			h.bot.Send(tgbotapi.NewMessage(chatID, "Пожалуйста, введите текстовое название категории."))
			return
		}

		input := usecase.CreateCategoryInput{
			TelegramID: uint32(update.Message.From.ID),
			Name:       update.Message.Text,
		}

		err := h.uc.CreateCategory(input)
		if err != nil {
			log.Printf("Failed to create category: %v", err)
			h.bot.Send(tgbotapi.NewMessage(chatID, "Ошибка при создании категории."))
		} else {
			log.Printf("Category %s created successfully", input.Name)
			h.bot.Send(tgbotapi.NewMessage(chatID, "Категория успешно создана!"))
		}

		// Сбрасываем состояние ожидания
		h.waitingFor.Delete(chatID)
		msg := tgbotapi.NewMessage(chatID, "Категория успешно создана!")
		h.bot.Send(msg)
	} else {
		// Запрашиваем у пользователя ввод названия категории
		msg := tgbotapi.NewMessage(chatID, "Введите название категории:")
		_, err := h.bot.Send(msg)
		if err != nil {
			log.Printf("Failed to send request message: %v", err)
		}

		// Устанавливаем состояние ожидания для этого чата
		h.waitingFor.Store(chatID, true)
	}
}
