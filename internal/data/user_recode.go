package data

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email"`
	ProfilePic FileRecode `gorm:"foreignkey:ID;",json:"ProfilePic"`
}
type UserLogin struct {
	gorm.Model
	Username  string
	Password  string
	UserRefer uint
	User      User `gorm:"foreignkey:UserRefer;association_autoupdate:false"`
}

type UserList []User
