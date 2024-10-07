package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetMainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	buttonRow := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("📄 Мои заметки"),
		tgbotapi.NewKeyboardButton("➕ Новая заметка"),
	}

	keyboard := tgbotapi.NewReplyKeyboard(
		buttonRow,
	)

	return keyboard
}
