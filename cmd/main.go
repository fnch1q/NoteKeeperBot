package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	botToken, err := token()
	if err != nil {
		handleError(err)
		return
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		handleError(err)
		return
	}

	startBot(bot)
}

func token() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		return "", fmt.Errorf("TELEGRAM_BOT_TOKEN is not set")
	}

	return botToken, nil
}

func startBot(bot *tgbotapi.BotAPI) {
	chatIDstr := os.Getenv("MY_CHAT_ID")
	chatID, err := strconv.ParseInt(chatIDstr, 10, 64)
	if err != nil {
		handleError(err)
		return
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	msg := tgbotapi.NewMessage(chatID, "Привет!")
	bot.Send(msg)

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			handleMessage(bot, update.Message)
		}
	}
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	bot.Send(msg)
}

func handleError(err error) {
	log.Printf("Error: %v\n", err)
}
