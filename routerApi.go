package main

import (
	"github.com/KhaleghiDev/GraphQLVotingSystem/app/controller/api"
	"github.com/gin-gonic/gin"
)

func setupApiRouter() *gin.Engine {
	r := gin.Default()
	//set defult address api/v1
	apis := r.Group("/api")
	v1 := apis.Group("/v1")

	//router start
	v1.GET("/dashbord", api.All)
	return r
}
