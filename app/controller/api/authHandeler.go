package api

import "github.com/gin-gonic/gin"

//Login User
func LoginUser(c *gin.Context) {
	
	user := c.Param("email")
	pass := c.Param("password")
	if user != nil {
		
	}

}

//Register User
func RegisterUser(c *gin.Context) {

}
