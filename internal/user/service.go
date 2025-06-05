package user

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var errInvalidCredentials = errors.New("invalid credentials")

type UserService interface {
	Register(user User) (*User, error)
	Login(phone, password string) (*User, error)
	Authorize(tokenString string) (*User, error)
	PhoneExists(phone string) (bool, error)
	EmailExists(email string) (bool, error)
	Update(id uint, updates map[string]interface{}) error
}

type GormUserService struct {
	repo UserRepo
}

func NewGormUserService(repo UserRepo) *GormUserService {
	return &GormUserService{repo: repo}
}

func (s *GormUserService) Register(newUser User) (*User, error) {
	var e error
	newUser.Password, e = HashPassword(newUser.Password)
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

	if !VerifyPassword(password, user.Password) {
		return nil, errInvalidCredentials
	}

	return user, nil
}

func (svc *GormUserService) Authorize(tokenString string) (*User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, errors.New("expired token")
	}

	return svc.repo.FindById(uint(claims["id"].(float64)))
}

func (svc *GormUserService) PhoneExists(phone string) (bool, error) {
	user, e := svc.repo.FindByPhone(phone)
	return user != nil, e
}

func (svc *GormUserService) EmailExists(email string) (bool, error) {
	user, e := svc.repo.FindByEmail(email)
	return user != nil, e
}

func (svc *GormUserService) Update(id uint, updates map[string]interface{}) error {
	user, err := svc.repo.FindById(id)
	if err != nil {
		return errInvalidCredentials
	}
	for field, value := range updates {
		switch field {
		case "full_name":
			user.FullName = value.(string)

		case "phone":
			user.Phone = value.(string)

		case "email":
			user.Email = value.(string)

		case "password":
			user.Password = value.(string)

		case "bank_info":
			user.BankInfo = value.(BankInfo)

		case "address":
			user.Address = value.(string)

		case "profileImageBase64":
			user.ProfileImageBase64 = value.(string)
		}
	}
	return svc.repo.Save(user)
}
