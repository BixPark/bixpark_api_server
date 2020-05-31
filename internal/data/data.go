package data

import "github.com/jinzhu/gorm"

func SetupDB(db *gorm.DB) {
	db.AutoMigrate(&User{}, &UserLogin{})
}
