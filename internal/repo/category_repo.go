package repo

import (
	"NoteKeeperBot/internal/entities"
	"time"

	"gorm.io/gorm"
)

type CategoryGORM struct {
	ID          uint32    `gorm:"primary_key"`
	UserID      uint32    `gorm:"column:user_id"`
	Name        string    `gorm:"column:category_name"`
	CreatedDate time.Time `gorm:"column:created_date"`
}

type CategoryDB struct {
	db        *gorm.DB
	tableName string
}

func NewCategoryDB(db *gorm.DB) CategoryDB {
	return CategoryDB{
		db:        db,
		tableName: "categories",
	}
}

func (u CategoryDB) Create(Category entities.Category) error {
	var CategoryGORM = CategoryGORM{
		UserID:      Category.GetUserID(),
		Name:        Category.GetName(),
		CreatedDate: Category.GetCreatedDate(),
	}

	if err := u.db.Table(u.tableName).Create(&CategoryGORM).Error; err != nil {
		return err
	}

	return nil
}

func (u CategoryDB) GetAll(userID uint32) ([]entities.Category, int64, error) {
	var categoriesGORM []CategoryGORM
	var total int64

	if err := u.db.Table(u.tableName).Where("user_id = ?", userID).Count(&total).Find(&categoriesGORM).Error; err != nil {
		return nil, 0, err
	}

	var categories []entities.Category
	for _, categoryGORM := range categoriesGORM {
		category := entities.NewCategory(
			categoryGORM.ID,
			categoryGORM.UserID,
			categoryGORM.Name,
			categoryGORM.CreatedDate,
		)
		categories = append(categories, category)
	}

	return categories, total, nil
}
