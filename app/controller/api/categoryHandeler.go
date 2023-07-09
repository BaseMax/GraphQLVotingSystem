package api

import (
	"github.com/KhaleghiDev/GraphQLVotingSystem/Config"
	"github.com/KhaleghiDev/GraphQLVotingSystem/app/models"
	"github.com/gin-gonic/gin"
)

// all data category
func AllCategory(c *gin.Context) {

	var category []models.Categories
	Config.DB.Find(&category)

	c.JSON(200, gin.H{
		"message": "all category",
		"data":    category,
	})

}

// show data category
func ShowCategory(c *gin.Context) {
	// get id
	id := c.Param("id")

	var category models.Categories
	Config.DB.First(&category, id)
	c.JSON(200, gin.H{
		"status":  true,
		"message": "show data id : " + id,
		"data":    category,
	})
}

// create save data to db
func CreateCategory(c *gin.Context) {
	var body struct {
		Title     string
		Parint_id uint
	}
	c.Bind(&body)

	movies := models.Categories{Title: body.Title, Parint_id: body.Parint_id}

	res := Config.DB.Create(&movies)
	if res != nil {
		c.Status(400)
		c.JSON(400, gin.H{
			"status": false,
		})
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "create data",
	})
}

// update data category
func UpdateCategory(c *gin.Context) {

	id := c.Param("id")

	var category models.Categories

	Config.DB.First(&category, id)
	var body struct {
		Title     string
		Parint_id uint
	}
	c.Bind(&body)

	res := Config.DB.Model(&category).Updates(models.Categories{
		Title:     body.Title,
		Parint_id: body.Parint_id,
	})
	if res != nil {
		c.Status(400)
		c.JSON(400, gin.H{
			"status": false,
		})
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "update data",
		"data":    category,
	})
}
