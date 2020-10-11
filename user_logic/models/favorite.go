package models

import "github.com/jinzhu/gorm"

type Favorite struct {
	gorm.Model
	UserID   int    `json:"UserID"`
	JsonText string `gorm:"type:text"`
}
