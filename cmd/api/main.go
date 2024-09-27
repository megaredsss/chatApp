package main

import (
	"chatApp/internal/database"
	jwtpackage "chatApp/internal/jwt"
	routes "chatApp/internal/server"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	database.ConnectToDb()
	r := gin.Default()
	r.LoadHTMLGlob("templates/html/*")
	routes.Login(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	jwtpackage.CreateSecretKey()
	jwtpackage.CreateToken("testUsername")
}
