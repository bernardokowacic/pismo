package controller

import (
	"math"
	"net/http"
	"pismo/entity"
	"pismo/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func InsertTransaction(transactionService service.TransactionServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Debug().Msg("end-point POST /transactions requested")

		var postData entity.Transaction
		err := ctx.ShouldBindJSON(&postData)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
			return
		}
		postData.Amount = math.Round(postData.Amount*100) / 100 // round to nearest

		response, err := transactionService.Insert(postData)
		if err != nil {
			log.Error().Msg(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Debug().Msg("end-point POST /transactions finished")

		ctx.JSON(http.StatusOK, response)
	}
}
