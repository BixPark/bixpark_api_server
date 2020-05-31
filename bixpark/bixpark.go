package bixpark

import (
	"bixpark_server/config"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	log.Println(app.Config.App.Name, " app_route :: ", path)
	return path
}

func (app *BixParkApp) Init() {
	app.Router.Use(cors)
	app.Router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(app.Config.App.Name, "Route Log :: ", r.URL.String(), " ", r.Method)
			handler.ServeHTTP(w, r)
			return
		})
	})

}
func (app *BixParkApp) Finalize() {
	spa := spaHandler{staticPath: "./spa_web/build", indexPath: "index.html"}
	app.Router.PathPrefix("/").Handler(spa)
	app.Router.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir(app.Config.App.MediaPath))))
}

func (app *BixParkApp) Log(tag string, messages ...interface{}) {
	prefix := make([]interface{}, 0)
	prefix = append(prefix, app.Config.App.Name, " ", tag, " :: ")
	messages = append(prefix, messages...)
	log.Println(messages...)
}

func openLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		log.SetOutput(lf)
	}
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

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path := r.URL.Path //filepath.Abs(r.URL.Path)
	//if err != nil {
	//	// if we failed to get the absolute path respond with a 400 bad request
	//	// and stop
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	//log.Println(h.staticPath, "  ", path, " ", r.URL.Path)

	// check whether a file exists at the given path
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
