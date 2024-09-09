package initializers

import (
	"example.com/m/v2/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB () *gorm.DB {

  DB, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  
  
  DB.AutoMigrate(&models.User{})  
  return DB
}
  