package controller

import (
	"net/http"
	"pismo/entity"
	"pismo/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func InsertAccount(accountService service.AccountServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Debug().Msg("end-point POST /accounts requested")

		var postData entity.Account
		err := ctx.ShouldBindJSON(&postData)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
			return
		}

		response, err := accountService.Insert(postData)
		if err != nil {
			log.Error().Msg(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Debug().Msg("end-point GET /accounts finished")

		ctx.JSON(http.StatusOK, response)
	}
}

func FindAccount(accountService service.AccountServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Debug().Msg("end-point GET /account requested")

		requestedAccount := struct {
			ID uint64 `uri:"accountId" binding:"required"`
		}{}
		err := ctx.ShouldBindUri(&requestedAccount)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
			return
		}

		response, err := accountService.Get(requestedAccount.ID)
		if err != nil {
			log.Error().Msg(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		log.Debug().Msg("end-point GET /account finished")

		ctx.JSON(http.StatusOK, response)
	}
}
