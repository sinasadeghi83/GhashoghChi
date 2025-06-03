package user

import "gorm.io/gorm"

type UserRepo interface {
	Create(newUser User) (User, error)
}

type GormUserRepo struct {
	db *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) *GormUserRepo {
	return &GormUserRepo{db: db}
}

func (repo *GormUserRepo) Create(newUser User) (User, error) {
	result := repo.db.Create(&newUser)
	return newUser, result.Error
}
