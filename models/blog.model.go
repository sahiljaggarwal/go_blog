package models

import "gorm.io/gorm"


type Blog struct {
	gorm.Model

	Title string `gorm:"not null"`
	Content string `gorm:"type:text"`
	UserId uint 
	User User `gorm:"foreignKey:UserId"`
	Image *string `gorm:"type:varchar(255)"`
}