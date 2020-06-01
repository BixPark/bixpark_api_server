package api

import (
	"bixpark_server/bixpark"
	"bixpark_server/internal/data"
	"bixpark_server/internal/results"
	"bixpark_server/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var TAG = "USER_API"

type UserApi struct {
	app *bixpark.BixParkApp
}

func (api UserApi) SetupRoutes() {
	content := "users"
	// User Login
	api.app.Router.HandleFunc(api.app.BuildPath("authenticate"), api.authenticate).Methods("POST")

	// Users CURD Operations
	api.app.Router.HandleFunc(api.app.BuildPath(content), api.post).Methods("POST")
	api.app.Router.HandleFunc(api.app.BuildPath(content, "{limit}", "{offset}"), api.getAll).Methods("GET")
	api.app.Router.HandleFunc(api.app.BuildPath(content, "{id}"), api.update).Methods("PUT")
	api.app.Router.HandleFunc(api.app.BuildPath(content, "{id}"), api.delete).Methods("DELETE")
	api.app.Router.HandleFunc(api.app.BuildPath(content, "{id}"), api.get).Methods("GET")
}

func (api UserApi) authenticate(w http.ResponseWriter, r *http.Request) {
	var loginCredentials data.Credentials
	err := json.NewDecoder(r.Body).Decode(&loginCredentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userService := service.UserService{
		DB: api.app.DB,
	}
	api.app.Log(TAG, loginCredentials.Password, " ", loginCredentials.Username)
	loginCredentials.Hash(api.app.Config)
	user, err := userService.Authenticate(&loginCredentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loginResponse := results.SingleResponse{
		Data:    user,
		Message: results.Message{},
	}

	json.NewEncoder(w).Encode(loginResponse)
}

// _/users - POST - SAVE
func (api UserApi) post(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(100000)

	api.app.Log(TAG, err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get a ref to the parsed multipart form
	m := r.MultipartForm
	profilePics := m.File["profilePic"]
	if len(profilePics) > 0 {
		profilePic, err := profilePics[0].Open()

		api.app.Log(TAG, profilePics[0].Filename)

		defer profilePic.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dst, err := os.Create(fmt.Sprintf("%s/%s", api.app.Config.App.MediaPath, profilePics[0].Filename))
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, profilePic); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	user := data.User{
		FirstName: r.FormValue("firstName"),
		LastName:  r.FormValue("lastName"),
		Email:     r.FormValue("email"),
		ProfilePic: data.FileRecode{
			Name: profilePics[0].Filename,
			Type: "",
			Path: profilePics[0].Filename,
		},
	}

	credentials := data.Credentials{
		Username: r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	credentials.Hash(api.app.Config)
	user.ProfilePic.Encode()
	userService := service.UserService{
		DB: api.app.DB,
	}
	userService.Save(&user, &credentials)

	response := results.SingleResponse{
		Data: &user,
		Message: results.Message{
			Status:  200,
			Message: "Success",
		},
	}
	json.NewEncoder(w).Encode(response)
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
