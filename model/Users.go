package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model `swaggerignore:"true"`
	Email   string `gorm:"unique;not null" json:"email"`
	Password   string `gorm:"not null" json:"password"`
}