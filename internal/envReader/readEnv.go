package envReader

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	dbName     string
	dbUsername string
	dbPassword string
	dbHost     string
	dbPort     string
}

var enviromment Env

func GetDbName(e Env) string {
	return e.dbName
}

func GetDbUsername(e Env) string {
	return e.dbUsername
}

func GetDbPassword(e Env) string {
	return e.dbPassword
}

func GetDbHost(e Env) string {
	return e.dbHost
}

func GetDbPort(e Env) string {
	return e.dbPort
}

func ReadEnv() Env {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	enviromment.dbName = os.Getenv("DB_DATABASE")
	enviromment.dbUsername = os.Getenv("DB_USERNAME")
	enviromment.dbPassword = os.Getenv("DB_PASSWORD")
	enviromment.dbHost = os.Getenv("DB_HOST")
	enviromment.dbPort = os.Getenv("DB_PORT")
	return enviromment
}
