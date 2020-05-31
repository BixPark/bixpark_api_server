package data

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
type UserLogin struct {
	Username  string
	Password  string
	UserRefer uint
	User      User `gorm:"foreignkey:UserRefer;association_autoupdate:false"`
}

type UserList []User
