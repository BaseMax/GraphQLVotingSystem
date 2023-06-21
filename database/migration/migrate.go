package main

import (
		"GraphQLVotingSystem/Config"
		"GraphQLVotingSystem/App/models"
)

func init() {
	Config.LoadEnvVarables()
	Config.ConectToDb()
}
func main() {
	Config.DB.AutoMigrate(&models.User{})

}
