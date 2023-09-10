package api

import (
	"fmt"

	"github.com/KhaleghiDev/GraphQLVotingSystem/Config"
	"github.com/KhaleghiDev/GraphQLVotingSystem/app/models"
	"github.com/gin-gonic/gin"
)

func init() {
	Config.LoadEnvVarables()
	Config.ConectToDb()
}

// Login User
func LoginUser(c *gin.Context) {

	username := c.Param("email")
	password := c.Param("password")
	if username == "" || password == "" {
		fmt.Println("نام کاربری و رمز عبور را وارد نکردید .")
	}
	var user models.User
	Config.DB.Where("Username:? AND Password:?", username, password).First(&user)
	c.JSON(200, gin.H{
		"status":  true,
		"message": "show data id : ",
		"data":    user,
	})
}

// Register User
func RegisterUser(c *gin.Context) {

}
