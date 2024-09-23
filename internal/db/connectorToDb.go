package db

import (
	"chatApp/internal/envReader"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() {
	env := envReader.ReadEnv()
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", envReader.GetDbHost(env), envReader.GetDbPort(env), envReader.GetDbUsername(env), envReader.GetDbPassword(env), envReader.GetDbName(env))), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database: %s", err)
	} else {
		fmt.Println(db.Migrator().CurrentDatabase())
	}
	runAutomigration(db)
}
