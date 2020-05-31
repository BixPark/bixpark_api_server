package main

import (
	"bixpark_server/api"
	"bixpark_server/bixpark"
	"bixpark_server/config"
	"bixpark_server/internal/data"
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	var wait time.Duration

	var configPath string
	flag.StringVar(&configPath, "config", os.Getenv("BIX_PARK_APP_CONFIG"), "App config locations")
	flag.Parse()

	config, err := config.NewConfig(configPath)
	if err != nil {
		println("Error on reading config")
	}
	// Load App configs
	log.Println("Load config success ", config.App.Name)

	//dB
	db, err := DBConnect(config)
	if err != nil {
		log.Println("Error in DB Connect")
	}
	defer db.Close()

	r := mux.NewRouter()

	app := bixpark.BixParkApp{
		Router: r,
		DB:     db,
		Config: config,
	}

	app.Router.HandleFunc("/api/v1", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Category: %v\n", vars["category"])
	})

	api.SetupRoutes(&app)
	data.SetupDB(&app)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      app.Router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

}
