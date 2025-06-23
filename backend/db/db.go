package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Url struct {
  gorm.Model
  Link string	`json:"link" validate:"required,min_length=5"`
  Validity int	`json:"validity"`
  ShortCode string	`json:"shortcode"`
  ExpireAt time.Time
  ShortLink string 
}

func GetDB()(*gorm.DB,error){
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    return nil,err 
  }
  if err := db.AutoMigrate(&Url{}); err != nil {
	return nil,err 
  }
  return db,nil 
}
