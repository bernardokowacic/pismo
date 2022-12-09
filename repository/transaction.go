package repository

import (
	"pismo/entity"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type TransactionInterface interface {
	Insert(transaction entity.Transaction) (entity.Transaction, error)
}

type transactionRepositoryStruct struct {
	DbConn *gorm.DB
}

func NewTransactionRepository(dbConn *gorm.DB) TransactionInterface {
	return &transactionRepositoryStruct{DbConn: dbConn}
}

func (t *transactionRepositoryStruct) Insert(transaction entity.Transaction) (entity.Transaction, error) {
	result := t.DbConn.Model(&entity.Transaction{}).Create(&transaction)
	if result.Error != nil {
		log.Error().Msg(result.Error.Error())
		return entity.Transaction{}, result.Error
	}

	return transaction, nil
}
