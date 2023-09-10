package main

import (
	"fmt"
	"net/http"

	"github.com/KhaleghiDev/GraphQLVotingSystem/app/controller/api"
	"github.com/gin-gonic/gin"
)

func setupApiRouter() *gin.Engine {
	r := gin.Default()
	//set defult address api/v1
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})

	
	apis := r.Group("/api")
	v1 := apis.Group("/v1")

	//router start
	v1.GET("/dashbord", api.All)

	//category api
	categories := v1.Group("/category")
	categories.GET("/", api.AllCategory)
	categories.GET("/:id", api.ShowCategory)
	categories.POST("/", api.CreateCategory)
	categories.POST("/:id", api.UpdateCategory)
	//users api
	auth := v1.Group("/users/auth")
	auth.POST("/login", api.LoginUser)
	auth.POST("/register", api.RegisterUser)

	return r
}
