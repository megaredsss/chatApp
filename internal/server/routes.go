package routes

import (
	"chatApp/internal/database"
	"chatApp/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(r *gin.Engine) {
	r.GET("/login", userLogin)
	r.POST("/login", createUser)
}

func signUp(c *gin.Context) {

}

func userLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
	fmt.Println(c.PostForm("email"))

}

func createUser(c *gin.Context) {
	user := models.User{
		Email:    c.PostForm("email"),
		Name:     c.PostForm("name"),
		Password: c.PostForm("password"),
	}
	database.AddUserToDb(user)
}
