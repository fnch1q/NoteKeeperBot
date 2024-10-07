package usecase

import (
	"NoteKeeperBot/internal/entities"
	"time"
)

type CreateCategoryInput struct {
	TelegramID uint32
	Name       string
}

type CreateCategoryUseCase struct {
	categoryRepo entities.CategoryRepository
	userRepo     entities.UserRepository
}

func NewCreateCategoryUseCase(
	categoryRepo entities.CategoryRepository,
	userRepo entities.UserRepository,
) CreateCategoryUseCase {
	return CreateCategoryUseCase{
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (uc CreateCategoryUseCase) CreateCategory(input CreateCategoryInput) error {
	user, err := uc.userRepo.GetByTelegramID(input.TelegramID)
	if err != nil {
		return err
	}

	category := entities.NewCategoryCreate(
		user.GetID(),
		input.Name,
		time.Now(),
	)
	err = uc.categoryRepo.Create(category)
	if err != nil {
		return err
	}
	return nil
}
