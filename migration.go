package users

import (
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
}
