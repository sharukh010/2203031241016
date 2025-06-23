package db

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Url struct {
  gorm.Model
  Link string
  Validity int
  ShortCode string
}

func init() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&Url{})
}
