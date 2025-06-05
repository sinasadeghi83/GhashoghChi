package user

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(newUser User) (*User, error)
	FindByPhone(phone string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindById(id uint) (*User, error)
	Save(user *User) error
}

type GormUserRepo struct {
	db *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) *GormUserRepo {
	return &GormUserRepo{db: db}
}

func (repo *GormUserRepo) Create(newUser User) (*User, error) {
	result := repo.db.Create(&newUser)
	return &newUser, result.Error
}

func (repo *GormUserRepo) FindByPhone(phone string) (*User, error) {
	var user User
	err := repo.db.Where(&User{Phone: phone}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *GormUserRepo) FindByEmail(email string) (*User, error) {
	var user User
	err := repo.db.Where(&User{Email: email}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *GormUserRepo) FindById(id uint) (*User, error) {
	var user User
	err := repo.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *GormUserRepo) Save(user *User) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		if err := tx.First(user, user.ID).Error; err != nil {
			return fmt.Errorf("failed to reload user: %w", err)
		}

		return nil
	})

	return err
}
