package user

import "gorm.io/gorm"

type UserRepo interface {
	Create(newUser User) (*User, error)
	FindByPhone(phone string) (*User, error)
	FindById(id uint) (*User, error)
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

func (repo *GormUserRepo) FindById(id uint) (*User, error) {
	var user User
	err := repo.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
