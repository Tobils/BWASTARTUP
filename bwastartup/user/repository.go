package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(id int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	// error
	if err != nil {
		return user, err
	}

	// berhasil
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error

	// error
	if err != nil {
		return user, err
	}

	// berhasil
	return user, nil
}

func (r *repository) FindByID(id int) (User, error) {
	var user User
	err := r.db.Where("id = ? ", id).Find(&user).Error

	// error
	if err != nil {
		return user, err
	}

	// berhasil
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	// error
	if err != nil {
		return user, err
	}

	// berhasil
	return user, nil
}
