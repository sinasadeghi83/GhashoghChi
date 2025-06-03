package user

import u "github.com/sinasadeghi83/ghashoghchi/internal/user"

type RegisterInput struct {
	FullName           string     `json:"full_name" binding:"required"`
	Phone              string     `json:"phone" binding:"required,min=10,max=15"`
	Email              string     `json:"email" binding:"required,email"`
	Password           string     `json:"password" binding:"required,min=4,max=255"`
	Role               u.RoleType `json:"role" binding:"required"`
	Address            string     `json:"address" binding:"required"`
	ProfileImageBase64 string     `json:"profileImageBase64"`
	BankInfo           u.BankInfo `json:"bank_info"`
}

type BankInfo struct {
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
}
