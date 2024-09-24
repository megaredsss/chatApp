package main

import (
	"chatApp/internal/db"
	jwtpackage "chatApp/internal/jwt"

	_ "github.com/lib/pq"
)

func main() {
	db.ConnectToDb()
	jwtpackage.CreateSecretKey()
	jwtpackage.CreateToken("testUsername")
}
