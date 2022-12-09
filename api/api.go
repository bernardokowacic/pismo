package api

import (
	"os"
	"pismo/routes"
	"pismo/service"

	"github.com/gin-gonic/gin"
)

// Start initializes Gin API
func Start(accountService service.AccountServiceInterface, transactionService service.TransactionServiceInterface) *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()

	routes.GetRoutes(router, accountService, transactionService)
	return router
}
