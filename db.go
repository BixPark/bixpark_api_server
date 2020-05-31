package main

import (
	"bixpark_server/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"net/url"
)

func DBConnect(config *config.Config) (db *gorm.DB, err error) {
	dsn := url.URL{
		User:     url.UserPassword(config.DB.User, config.DB.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", config.DB.Host, config.DB.Port),
		Path:     config.DB.BbName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	log.Println(dsn.String())

	db, err = gorm.Open("postgres", dsn.String())
	db.LogMode(true)

	return
}
