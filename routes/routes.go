package routes

import (
	"pismo/controller"
	"pismo/service/account"

	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine, AccountService account.AccountServiceInterface) {
	router.POST("/accounts", controller.InsertAccount(AccountService))
	router.GET("/account/:accountId", controller.FindAccount(AccountService))
	// router.POST("/Transactions", controller.Insert(AccountService))
}
