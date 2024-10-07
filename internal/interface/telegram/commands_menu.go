package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetCommandMenu(bot *tgbotapi.BotAPI) error {
	commands := []tgbotapi.BotCommand{
		// {
		// 	Command:     "start",
		// 	Description: "Начало работы с ботом",
		// },
		{
			Command:     "add_category",
			Description: "Добавить категорию",
		},
		{
			Command:     "delete_category",
			Description: "Удалить категорию",
		},
		{
			Command:     "add_tag",
			Description: "Добавить тег",
		},
		{
			Command:     "help",
			Description: "Помощь",
		},
	}

	config := tgbotapi.NewSetMyCommands(commands...)

	// Устанавливаем команды для бота
	_, err := bot.Request(config)
	if err != nil {
		log.Printf("Error setting command menu: %v", err)
		return err
	}

	return nil
}
