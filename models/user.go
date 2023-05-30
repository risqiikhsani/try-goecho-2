package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Account     Account
	Name        string `form:"name" json:"name" xml:"name" query:"name"`
	Gender      string `form:"gender" json:"gender" xml:"gender" query:"gender"`
	DateOfBirth string `form:"dateOfBirth" json:"dateOfBirth" xml:"dateOfBirth" query:"dateOfBirth"`
	About       string `form:"about" json:"about" xml:"about" query:"about"`
	Country     string `form:"country" json:"country" xml:"country" query:"country"`
	Verified    bool   `form:"verified" json:"verified" xml:"verified" query:"verified"`
}
