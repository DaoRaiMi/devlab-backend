package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
	RoleId   uint   `gorm:"column:role_id"`
	Status   uint16 `gorm:"column:status"`
}

func (User) TableName() string {
	return "user"
}

type UserList []*User
