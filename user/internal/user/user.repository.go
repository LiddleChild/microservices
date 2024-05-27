package user

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) GetAllUsers(results *[]User) error {
	return r.db.Model(&User{}).Find(&results).Error
}

func (r *Repository) CreateUser(u *User) error {
	return r.db.Model(&User{}).Create(u).Error
}
