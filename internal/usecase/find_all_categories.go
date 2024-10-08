package usecase

import (
	"NoteKeeperBot/internal/entities"
)

type FindAllCategoriesInput struct {
	TelegramID uint32
}

type FindAllCategoriesUsecase struct {
	categoryRepo entities.CategoryRepository
	userRepo     entities.UserRepository
}

func NewFindAllCategoriesUseCase(categoryRepo entities.CategoryRepository, userRepo entities.UserRepository) FindAllCategoriesUsecase {
	return FindAllCategoriesUsecase{
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (uc FindAllCategoriesUsecase) FindAllCategories(input FindAllCategoriesInput) ([]entities.Category, int64, error) {
	user, err := uc.userRepo.GetByTelegramID(input.TelegramID)
	if err != nil {
		return nil, 0, err
	}

	categories, total, err := uc.categoryRepo.FindAll(user.GetID())
	if err != nil {
		return nil, 0, err
	}

	return categories, total, err
}
