package migration

import (
	"github.com/KhaleghiDev/GraphQLVotingSystem/Config"
	"github.com/KhaleghiDev/GraphQLVotingSystem/app/models"
)

func init() {
	Config.LoadEnvVarables()
	Config.ConectToDb()
}
func main() {
	Config.DB.AutoMigrate(&models.Users{})

}
