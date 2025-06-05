package user

import "gorm.io/gorm"

type RoleType string

const (
	Buyer   RoleType = "buyer"
	Seller  RoleType = "seller"
	Courier RoleType = "courier"
)

type User struct {
	gorm.Model
	FullName           string   `gorm:"full_name" json:"full_name"`
	Phone              string   `gorm:"phone,unique" json:"phone"`
	Email              string   `gorm:"email,unique" json:"email"`
	Password           string   `gorm:"password" json:"-"`
	Role               RoleType `gorm:"role" json:"role"`
	Address            string   `gorm:"address" json:"address"`
	ProfileImageBase64 string   `gorm:"column:profileImageBase64" json:"profileImageBase64,omitempty"`
	BankInfo           BankInfo `gorm:"embedded" json:"bank_info,omitempty"`
}

type BankInfo struct {
	BankName      string `gorm:"bank_name" json:"bank_name,omitempty"`
	AccountNumber string `gorm:"account_number" json:"account_number,omitempty"`
}

func (User) TableName() string {
	return "users"
}
