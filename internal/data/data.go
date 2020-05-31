package data

import (
	"bixpark_server/bixpark"
	"log"
)

func SetupDB(app *bixpark.BixParkApp) {
	app.DB.AutoMigrate(&User{}, &UserLogin{})
	log.Println(app.Config.App.Name, " DB Initiate :: ", "User")
	log.Println(app.Config.App.Name, " DB Initiate :: ", "UserLogin")
}
