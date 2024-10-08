package telegram

import (
	"NoteKeeperBot/internal/usecase"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler struct {
	bot                      *tgbotapi.BotAPI
	createUserHandler        CreateUserHandler
	createCategoryHandler    CreateCategoryHandler
	deleteCategoryHandler    DeleteCategoryHandler
	findAllCategoriesHandler FindAllCategoriesHandler
	userComands              sync.Map
}

func NewMessageHandler(
	bot *tgbotapi.BotAPI,
	createUserUC usecase.CreateUserUseCase,
	createCategoryUC usecase.CreateCategoryUseCase,
	deleteCategoryUC usecase.DeleteCategoryUseCase,
	findAllCategoryUC usecase.FindAllCategoriesUsecase,
) MessageHandler {
	createUserHandler := NewCreateUserHandler(createUserUC, bot)
	createCategoryHandler := NewCreateCategoryHandler(createCategoryUC, bot)
	deleteCategoryHandler := NewDeleteCategoryHandler(deleteCategoryUC, bot)
	findAllCategoryHandler := NewFindAllCategoriesHandler(findAllCategoryUC, bot)
	return MessageHandler{
		bot:                      bot,
		createUserHandler:        createUserHandler,
		createCategoryHandler:    createCategoryHandler,
		deleteCategoryHandler:    deleteCategoryHandler,
		findAllCategoriesHandler: findAllCategoryHandler,
	}
}

func (mh *MessageHandler) HandleMessage(update tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.IsCommand() {
			mh.handleCommand(update)
		} else {
			mh.handleTextMessage(update)
		}
	}
}

func (mh *MessageHandler) handleCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		// Логика для команды /start
		mh.createUserHandler.Handle(update)
	case "add_category":
		// Сохраняем команду для пользователя
		mh.userComands.Store(update.Message.From.ID, "add_category")
		mh.bot.Send(tgbotapi.NewMessage(update.Message.From.ID, "Пожалуйста, введите текстовое название категории."))
	case "delete_category":
		mh.userComands.Store(update.Message.From.ID, "delete_category")
		mh.bot.Send(tgbotapi.NewMessage(update.Message.From.ID, "Пожалуйста, введите текстовое название категории."))
	}
}

func (mh *MessageHandler) handleTextMessage(update tgbotapi.Update) {
	lastCommand, ok := mh.userComands.Load(update.Message.From.ID)
	if update.Message.Text != "" {
		if ok {
			switch lastCommand.(string) {
			case "add_category":
				mh.createCategoryHandler.Handle(update)
				mh.userComands.Delete(update.Message.From.ID)
			case "delete_category":
				mh.deleteCategoryHandler.Handle(update)
				mh.userComands.Delete(update.Message.From.ID)
			}
		} else {
			switch update.Message.Text {
			case "📄 Мои заметки":
				mh.findAllCategoriesHandler.Handle(update)
			}
		}
	}
}
