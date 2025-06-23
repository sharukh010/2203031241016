package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Link struct {
  gorm.Model
  Url string	`json:"url"`
  Validity int	`json:"validity"`
  ShortCode string	`json:"shortcode"`
  ClickCount int	`json:"click_count"`
  ExpireAt time.Time
  ShortLink string 
}

func GetDB()(*gorm.DB,error){
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    return nil,err 
  }
  if err := db.AutoMigrate(&Link{}); err != nil {
	return nil,err 
  }
  return db,nil 
}
