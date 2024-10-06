package entities

import (
	"errors"
	"time"
)

type Category struct {
	id          uint32
	user_id     uint32
	name        string
	createdDate time.Time
}

type CategoryRepository interface {
	Create(Category Category) error
	GetAll(user_id uint32) ([]Category, int64, error)
}

var (
	ErrCategoryNotFound = errors.New("Category not found")
)

func NewCategory(
	id uint32,
	user_id uint32,
	name string,
	createdDate time.Time,
) Category {
	return Category{
		id:          id,
		user_id:     user_id,
		name:        name,
		createdDate: createdDate,
	}
}

func NewCategoryCreate(
	user_id uint32,
	name string,
	createdDate time.Time,
) Category {
	return Category{
		user_id:     user_id,
		name:        name,
		createdDate: createdDate,
	}
}

func (u Category) GetID() uint32 {
	return u.id
}

func (u Category) GetUserID() uint32 {
	return u.user_id
}

func (u Category) GetName() string {
	return u.name
}

func (u Category) GetCreatedDate() time.Time {
	return u.createdDate
}
