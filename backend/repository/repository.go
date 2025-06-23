package repository

import (
	"log"

	"github.com/sharukh010/url-shortner/db"
	"gorm.io/gorm"
)

var mydb *gorm.DB


func init(){
	db,err := db.GetDB()
	if err != nil {
		log.Panic("Failed to connect with db")
	}
	log.Println("Successfully connected to db")
	mydb = db 
}

func CreateLink(newLink *db.Url) error {
	if err := mydb.Create(newLink).Error; err != nil {
		log.Println("Error while creating new link:", err)
		return err
	}
	log.Println("New link created successfully")
	return nil
}