package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(r *gin.Engine) {
	r.GET("/login", userLogin)
	r.POST("/login", userLogin)
}

func signUp(c *gin.Context) {

}

func userLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
	fmt.Println(c.PostForm("email"))

}
