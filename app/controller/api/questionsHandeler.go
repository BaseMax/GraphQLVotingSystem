package api

import "github.com/gin-gonic/gin"

//Login User
func list(c *gin.Context) {

	code := c.Param("code")

	if code != "" {
		c.Status(400)
		c.JSON(400, gin.H{
			"status":  false,
			"massage": "مقادیر به درستی ارسال نشد ",
		})
	}

}

// show question
func ShowQuestions(c *gin.Context) {

	code := c.Param("code")

	if code != "" {
		c.Status(400)
		c.JSON(400, gin.H{
			"status":  false,
			"massage": "مقادیر به درستی ارسال نشد ",
		})
	}

}
