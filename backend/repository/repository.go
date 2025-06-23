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

func CreateLink(newLink *db.Link) error {
	if err := mydb.Create(newLink).Error; err != nil {
		return err
	}
	return nil
}
func GetLinkByShortCode(shortCode string) (*db.Link, error) {
	var link db.Link
	if err := mydb.Where("short_code = ?", shortCode).First(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

func IsLinkExists(shortCode string) bool {
	var count int64
	if err := mydb.Model(&db.Link{}).Where("short_code = ?", shortCode).Count(&count).Error; err != nil {
		log.Println("Error checking if link exists:", err)
		return false
	}
	return count > 0
}

func UpdateClickCount(shortCode string) error {
	var link db.Link
	if err := mydb.Where("short_code = ?", shortCode).First(&link).Error; err != nil {
		return err
	}
	link.ClickCount++
	if err := mydb.Save(&link).Error; err != nil {
		return err
	}
	return nil
}