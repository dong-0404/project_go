package model

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name     string  `gorm:"column:name" json:"name"`
	Email    string  `gorm:"column:email" json:"email"`
	Password string  `gorm:"column:password" json:"password"`
	RoleId   int     `gorm:"column:role_id" json:"role_id"`
	Role     []*Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	Id   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
