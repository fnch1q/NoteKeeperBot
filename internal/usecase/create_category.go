package usecase

// import (
// 	"NoteKeeperBot/internal/entities"
// 	"log"
// 	"time"
// )

// type CreateCategoryInput struct {
// 	TelegramID uint32
// 	Name       string
// }

// type CreateCategoryUseCase struct {
// 	categoryRepo entities.CategoryRepository
// 	userRepo     entities.UserRepository
// }

// func NewCreateCategoryUseCase(
// 	categoryRepo entities.CategoryRepository,
// 	userRepo entities.UserRepository,
// ) CreateCategoryUseCase {
// 	return CreateCategoryUseCase{
// 		categoryRepo: categoryRepo,
// 		userRepo:     userRepo,
// 	}
// }

// func (uc CreateCategoryUseCase) CreateCategory(input CreateCategoryInput) error {
// 	category := entities.NewCategoryCreate(
// 		input.UserID,
// 		input.Name,
// 		time.Now(),
// 	)
// 	err := uc.CategoryRepo.Create(Category)
// 	if err != nil {
// 		log.Printf("Failed to create Category: %v", err)
// 		return err
// 	}

// 	log.Printf("Category %s created successfully", category.GetName())
// 	return nil
// }
