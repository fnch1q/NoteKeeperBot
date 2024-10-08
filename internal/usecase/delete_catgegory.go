package usecase

import (
	"NoteKeeperBot/internal/entities"
)

type DeleteCategoryInput struct {
	TelegramID uint32
	Name       string
}

type DeleteCategoryUseCase struct {
	categoryRepo entities.CategoryRepository
	userRepo     entities.UserRepository
}

func NewDeleteCategoryUseCase(
	categoryRepo entities.CategoryRepository,
	userRepo entities.UserRepository,
) DeleteCategoryUseCase {
	return DeleteCategoryUseCase{
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (uc DeleteCategoryUseCase) DeleteCategory(input DeleteCategoryInput) error {
	user, err := uc.userRepo.GetByTelegramID(input.TelegramID)
	if err != nil {
		return err
	}

	category, err := uc.categoryRepo.FindByName(user.GetID(), input.Name)
	if err != nil {
		return err
	}

	err = uc.categoryRepo.Delete(category.GetID())
	if err != nil {
		return err
	}

	return nil
}
