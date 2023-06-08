package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Insert(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}
