package dto

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type CreateUser struct {
	*gorm.Model
	jwt.StandardClaims
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

func (CreateUser) TableName() string {
	return "users"
}
