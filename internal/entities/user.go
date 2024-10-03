package entities

import (
	"errors"
	"time"
)

type User struct {
	id          uint32
	telegramID  uint32
	name        string
	createdDate time.Time
}

type UserRepository interface {
	Create(user User) error
	GetByTelegramID(telegramID uint32) (User, error)
}

var (
	ErrUserNotFound = errors.New("user not found")
)

func NewUser(
	id uint32,
	telegramID uint32,
	name string,
	createdDate time.Time,
) User {
	return User{
		id:          id,
		telegramID:  telegramID,
		name:        name,
		createdDate: createdDate,
	}
}

func NewUserCreate(
	telegramID uint32,
	name string,
	createdDate time.Time,
) User {
	return User{
		telegramID:  telegramID,
		name:        name,
		createdDate: createdDate,
	}
}

func (u User) GetID() uint32 {
	return u.id
}

func (u User) GetTelegramID() uint32 {
	return u.telegramID
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetCreatedDate() time.Time {
	return u.createdDate
}
