package routes

import (
	"pismo/controller"
	"pismo/service"

	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine, accountService service.AccountServiceInterface, transactionService service.TransactionServiceInterface) {
	router.POST("/accounts", controller.InsertAccount(accountService))
	router.GET("/account/:accountId", controller.FindAccount(accountService))
	router.POST("/transactions", controller.InsertTransaction(transactionService))
}
