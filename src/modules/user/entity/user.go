package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int32
	Name     string
	Email    string
	Password string
}
