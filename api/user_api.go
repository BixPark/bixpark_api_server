package api

import (
	"bixpark_server/bixpark"
	"bixpark_server/internal/data"
	"bixpark_server/internal/results"
	"bixpark_server/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type UserApi struct {
	app *bixpark.BixParkApp
}

func (api UserApi) SetupRoutes() {
	content := "users"
	// Users
	api.app.Router.HandleFunc(api.app.BuildPath(content), api.post).Methods("POST")
	api.app.Router.HandleFunc(api.app.BuildPath(content, "{limit}", "{offset}"), api.getAll).Methods("GET")
	api.app.Router.HandleFunc(api.app.BuildPath(content, "{id}"), api.update).Methods("PUT")
	api.app.Router.HandleFunc(api.app.BuildPath(content, "{id}"), api.delete).Methods("DELETE")
	api.app.Router.HandleFunc(api.app.BuildPath(content, "{id}"), api.get).Methods("GET")
}

// _/users/ - GET - GET ALL
func (api UserApi) getAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit, _ := strconv.ParseInt(vars["limit"], 10, 0)
	offset, _ := strconv.ParseInt(vars["offset"], 10, 0)
	log.Println(limit, offset)
	//

	tagService := service.UserService{
		DB: api.app.DB,
	}

	userList := tagService.GetAll(int(limit), int(offset))
	response := results.ListResponse{
		Data: userList,
		Message: results.Message{
			Status:  200,
			Message: "SUCCESS",
		},
		TotalCount: len(userList),
		Limit:      int(limit),
		Offset:     int(offset),
	}

	log.Println(response)
	json.NewEncoder(w).Encode(response)

}

// _/users - POST - SAVE
func (api UserApi) post(w http.ResponseWriter, r *http.Request) {
	var user data.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(user.FirstName)
	userService := service.UserService{
		DB: api.app.DB,
	}
	userService.Save(&user)

	response := results.SingleResponse{
		Data: &user,
		Message: results.Message{
			Status:  200,
			Message: "Success",
		},
	}
	json.NewEncoder(w).Encode(response)
}

// _/users - PUT - UPDATE
func (api UserApi) update(w http.ResponseWriter, r *http.Request) {
	var user data.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userService := service.UserService{
		DB: api.app.DB,
	}
	userService.Update(&user)

	response := results.SingleResponse{
		Data: &user,
		Message: results.Message{
			Status:  200,
			Message: "Success",
		},
	}
	json.NewEncoder(w).Encode(response)
}

// _/users/{id} - GET
func (api UserApi) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, _ := strconv.ParseInt(vars["id"], 10, 0)

	tagService := service.UserService{
		DB: api.app.DB,
	}
	user := tagService.Get(int(ID))

	response := results.SingleResponse{
		Data: user,
		Message: results.Message{
			Status:  200,
			Message: "Success",
		},
	}

	json.NewEncoder(w).Encode(response)
}

// _/users/{id} - DELETE
func (api UserApi) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, _ := strconv.ParseInt(vars["id"], 10, 0)

	tagService := service.UserService{
		DB: api.app.DB,
	}
	tagService.Delete(int(ID))

	response := results.SingleResponse{
		Message: results.Message{
			Status:  200,
			Message: "Success",
		},
	}
	json.NewEncoder(w).Encode(response)
}
