package models

import "github.com/jinzhu/gorm"

type History struct {
	gorm.Model
	UserID   int    `json:"UserID"`
	JsonText string `gorm:"type:text"`
	Picture  string `gorm:"type:text"`
}
