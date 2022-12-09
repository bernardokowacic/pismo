package service

import (
	"pismo/entity"
	"pismo/repository"

	"github.com/rs/zerolog/log"
)

type TransactionServiceInterface interface {
	Insert(transaction entity.Transaction) (entity.Transaction, error)
}

type transactiontService struct {
	TransactionRepository repository.TransactionInterface
	AccountService        AccountServiceInterface
}

func NewTransactionService(transactionRepository repository.TransactionInterface, accountService AccountServiceInterface) TransactionServiceInterface {
	return &transactiontService{
		TransactionRepository: transactionRepository,
		AccountService:        accountService,
	}
}

func (t *transactiontService) Insert(transaction entity.Transaction) (entity.Transaction, error) {
	log.Debug().Msg("Creating new transaction")

	response, err := t.TransactionRepository.Insert(transaction)
	if err != nil {
		return entity.Transaction{}, err
	}

	return response, nil
}
