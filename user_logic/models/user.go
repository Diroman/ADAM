package models

import (
	"github.com/jinzhu/gorm"
)


type UserReg struct {
	gorm.Model
	Name          string
	LastName      string
	MiddleName    string
	CodeCountry   string
	BirthPlace    string
	birthDateTime string
	Phone         string
	Email         string `gorm:"type:varchar(100);unique_index"`
	Gender        string `json:"Gender"`
	Password      string `json:"Password"`
}

type UserLog struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `json:"Password"`
}
