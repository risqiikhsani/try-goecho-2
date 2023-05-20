package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name  string `form:"name" json:"name" xml:"name" query:"name"`
	Email string `form:"email" json:"email" xml:"email" query:"email"`
}
