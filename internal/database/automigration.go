package database

import (
	"chatApp/internal/models"
	"fmt"

	"gorm.io/gorm"
)

func runAutomigration(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Error during migration")
		return
	}
	if !db.Migrator().HasTable(&models.User{}) || !db.Migrator().HasColumn(&models.User{}, "email") || !db.Migrator().HasColumn(&models.User{}, "name") || !db.Migrator().HasColumn(&models.User{}, "password") {
		fmt.Println("Error!: Table or column doesn't exist")
		return
	}
}
