package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var errInvalidCredentials = errors.New("invalid credentials")

type UserService interface {
	Register(user User) (*User, error)
	Login(phone, password string) (*User, error)
}

type GormUserService struct {
	repo UserRepo
}

func NewGormUserService(repo UserRepo) *GormUserService {
	return &GormUserService{repo: repo}
}

func (s *GormUserService) Register(newUser User) (*User, error) {
	var e error
	newUser.Password, e = hashPassword(newUser.Password)
	if e != nil {
		return nil, e
	}
	return s.repo.Create(newUser)
}

func (s *GormUserService) Login(phone, password string) (*User, error) {
	user, err := s.repo.FindByPhone(phone)
	if err != nil {
		return nil, errInvalidCredentials
	}

	if !verifyPassword(password, user.Password) {
		return nil, errInvalidCredentials
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
