package controller

import (
	"net/http"
	"pismo/entity"
	"pismo/service/account"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func InsertAccount(AccountService account.AccountServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Debug().Msg("end-point POST /accounts requested")

		var postData entity.Account
		err := ctx.ShouldBindJSON(&postData)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
			return
		}

		response, err := AccountService.Insert(postData)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Debug().Msg("end-point GET /accounts finished")

		ctx.JSON(http.StatusOK, gin.H{"id": response.ID, "document_number": response.DocumentNumber})
	}
}
