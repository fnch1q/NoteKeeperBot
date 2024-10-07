package usecase

import (
	"NoteKeeperBot/internal/entities"
)

type FindAllCategoriesInput struct {
	UserID uint32
}

type FindAllCategoriesUsecase struct {
	categoryRepo entities.CategoryRepository
}

func NewFindAllCategoriesUseCase(categoryRepo entities.CategoryRepository) FindAllCategoriesUsecase {
	return FindAllCategoriesUsecase{
		categoryRepo: categoryRepo,
	}
}

func (uc FindAllCategoriesUsecase) FindAllCategories(input FindAllCategoriesInput) ([]entities.Category, int64, error) {
	categories, total, err := uc.categoryRepo.GetAll(input.UserID)
	if err != nil {
		return nil, 0, err
	}

	return categories, total, err
}
