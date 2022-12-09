package service

import (
	"errors"
	"pismo/entity"
	"pismo/repository"

	"github.com/rs/zerolog/log"
)

type TransactionServiceInterface interface {
	Insert(transaction entity.Transaction) (entity.Transaction, error)
}

type transactiontService struct {
	TransactionRepository   repository.TransactionInterface
	OperationTypeRepository repository.OperationTypeInterface
	accountRepository       repository.AccountInterface
}

func NewTransactionService(
	transactionRepository repository.TransactionInterface,
	operationTypeRepository repository.OperationTypeInterface,
	accountRepository repository.AccountInterface,
) TransactionServiceInterface {
	return &transactiontService{
		TransactionRepository:   transactionRepository,
		accountRepository:       accountRepository,
		OperationTypeRepository: operationTypeRepository,
	}
}

func (t *transactiontService) Insert(transaction entity.Transaction) (entity.Transaction, error) {
	log.Debug().Msg("Creating new transaction")

	_, err := t.accountRepository.Find(transaction.AccountID)
	if err != nil {
		log.Warn().Msg(err.Error())
		return entity.Transaction{}, errors.New("account not found")
	}

	_, err = t.OperationTypeRepository.Find(transaction.OperationTypeID)
	if err != nil {
		log.Warn().Msg(err.Error())
		return entity.Transaction{}, errors.New("operation type not found")
	}

	response, err := t.TransactionRepository.Insert(transaction)
	if err != nil {
		log.Warn().Msg(err.Error())
		return entity.Transaction{}, err
	}

	return response, nil
}
