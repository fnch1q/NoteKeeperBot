package usecase

import (
	"NoteKeeperBot/internal/entities"
	"log"
	"time"
)

type CreateUserInput struct {
	TelegramID uint32
	Name       string
}

type CreateUserUseCase struct {
	userRepo entities.UserRepository
}

func NewCreateUserUseCase(userRepo entities.UserRepository) CreateUserUseCase {
	return CreateUserUseCase{
		userRepo: userRepo,
	}
}

func (uc CreateUserUseCase) CreateUser(input CreateUserInput) error {
	_, err := uc.userRepo.GetByTelegramID(input.TelegramID)

	switch err {
	case entities.ErrUserNotFound:
		user := entities.NewUserCreate(
			input.TelegramID,
			input.Name,
			time.Now(),
		)
		err := uc.userRepo.Create(user)
		if err != nil {
			log.Printf("Failed to create user: %v", err)
			return err
		}

		log.Printf("User %s created successfully", user.GetName())
		return nil

	case nil:
		log.Printf("User %d already exists", input.TelegramID)
		return nil

	default:
		log.Printf("Failed to get user: %v", err)
		return err
	}
}
