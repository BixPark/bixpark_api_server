package api

import "bixpark_server/bixpark"

func SetupRoutes(app *bixpark.BixParkApp) {

	userApi := UserApi{app: app}
	userApi.SetupRoutes()
}
