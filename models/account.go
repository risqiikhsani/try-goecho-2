package models

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	UserID       uint
	Username     string `form:"username" json:"username" xml:"username" query:"username"`
	Password     string `form:"password" json:"password" xml:"password" query:"password"`
	Email        string `form:"email" json:"email" xml:"email" query:"email"`
	Phone_number int    `form:"phone_number" json:"phone_number" xml:"phone_number" query:"phone_number"`
	Verified     bool   `form:"verified" json:"verified" xml:"verified" query:"verified"`
}
