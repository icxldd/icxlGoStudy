package Entitys

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name string
	PassWord string
}