package repo

import (
	"NoteKeeperBot/internal/entities"
	"time"

	"gorm.io/gorm"
)

type userGORM struct {
	ID          uint32    `gorm:"primary_key"`
	TelegramID  uint32    `gorm:"column:telegram_id"`
	Name        string    `gorm:"column:username"`
	CreatedDate time.Time `gorm:"column:created_date"`
}

type UserDB struct {
	db        *gorm.DB
	tableName string
}

func NewUserDB(db *gorm.DB) UserDB {
	return UserDB{
		db:        db,
		tableName: "users",
	}
}

func (u UserDB) Create(user entities.User) error {
	var userGORM = userGORM{
		TelegramID:  user.GetTelegramID(),
		Name:        user.GetName(),
		CreatedDate: user.GetCreatedDate(),
	}

	if err := u.db.Table(u.tableName).Create(&userGORM).Error; err != nil {
		return err
	}

	return nil
}

func (u UserDB) GetByTelegramID(telegramID uint32) (entities.User, error) {
	var userGORM userGORM
	if err := u.db.Table(u.tableName).Where("telegram_id = ?", telegramID).First(&userGORM).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.User{}, entities.ErrUserNotFound
		}
		return entities.User{}, err
	}

	return entities.NewUser(
		userGORM.ID,
		userGORM.TelegramID,
		userGORM.Name,
		userGORM.CreatedDate,
	), nil
}
