package data

import (
	"github.com/jinzhu/gorm"
	"log"
)

type FileRecode struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
}

func (file FileRecode) Decode() {
	log.Println("Decode File ", file.Name)
}

func (file FileRecode) Encode() {
	log.Println("Encode File ", file.Name)
}
