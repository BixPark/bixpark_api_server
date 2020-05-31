package bixpark

import (
	"bixpark_server/config"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type BixParkApp struct {
	Router *mux.Router
	DB     *gorm.DB
	Config *config.Config
}

func (app *BixParkApp) BuildPath(content string, pathParams ...string) string {
	path := fmt.Sprintf("/%s/%s/%s", app.Config.Route.Prefix, app.Config.Route.Version, content)
	for _, pathParam := range pathParams {
		path = fmt.Sprintf("%s/%s", path, pathParam)
	}
	log.Println(path)
	return path
}

// CORS Middleware
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
		return
	})
}
