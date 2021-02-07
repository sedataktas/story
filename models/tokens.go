package models

import "github.com/jinzhu/gorm"

// Token represents taken table
type Token struct {
	gorm.Model
	AppID int
	Token string `gorm:"index;type:text"`
}
