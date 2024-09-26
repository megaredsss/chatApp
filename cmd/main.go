package main

import (
	"chatApp/internal/db"
	jwtpackage "chatApp/internal/jwt"
	"chatApp/internal/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db.ConnectToDb()
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		routes.Login(c)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	jwtpackage.CreateSecretKey()
	jwtpackage.CreateToken("testUsername")
}
