package data

import (
	"bixpark_server/bixpark"
)

func SetupDB(app *bixpark.BixParkApp) {
	app.DB.AutoMigrate(&User{}, &UserLogin{})
}
