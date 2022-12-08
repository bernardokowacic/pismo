package api

import (
	"os"
	"pismo/routes"
	"pismo/service/account"

	"github.com/gin-gonic/gin"
)

// Start initializes Gin API
func Start(accountService account.AccountServiceInterface) *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()

	routes.GetRoutes(router, accountService)
	return router
}
