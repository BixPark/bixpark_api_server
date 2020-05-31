## CREATE DATA CURD API

```go
package api

import (
	"bixpark_server/bixpark"
	"net/http"
)

type UserApi struct {
	app *bixpark.BixParkApp
}

func (api UserApi) SetupRoutes() {
	content := "users"
	// Users
	api.app.Router.HandleFunc(api.app.BuildPath(content), api.post).Methods("POST")
	api.app.Router.HandleFunc(api.app.BuildPath(content), api.getAll).Methods("GET")
	api.app.Router.HandleFunc(api.app.BuildPath(content), api.update).Methods("PUT")
	api.app.Router.HandleFunc(api.app.BuildPath(content), api.delete).Methods("DELETE")
	api.app.Router.HandleFunc(api.app.BuildPath(content), api.get).Methods("GET")
}

// _/users/ - GET - GET ALL
func (api UserApi) getAll(w http.ResponseWriter, r *http.Request) {}

// _/users - POST - SAVE
func (api UserApi) post(w http.ResponseWriter, r *http.Request) {}

// _/users - PUT - UPDATE
func (api UserApi) update(w http.ResponseWriter, r *http.Request) {}

// _/users/{id} - GET
func (api UserApi) get(w http.ResponseWriter, r *http.Request) {}

// _/users/{id} - DELETE
func (api UserApi) delete(w http.ResponseWriter, r *http.Request) {}

```