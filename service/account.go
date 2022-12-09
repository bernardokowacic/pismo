package service

import (
	"pismo/entity"
	"pismo/repository"

	"github.com/rs/zerolog/log"
)

type AccountServiceInterface interface {
	Insert(account entity.Account) (entity.Account, error)
	Get(accountID uint64) (entity.Account, error)
}

type AccountService struct {
	AccountRepository repository.AccountInterface
}

func NewService(accountRepository repository.AccountInterface) AccountServiceInterface {
	return &AccountService{
		AccountRepository: accountRepository,
	}
}

func (a *AccountService) Insert(account entity.Account) (entity.Account, error) {
	log.Debug().Msg("Creating new account")

	response, err := a.AccountRepository.Insert(account)
	if err != nil {
		return entity.Account{}, err
	}

	return response, nil
}

func (a *AccountService) Get(accountID uint64) (entity.Account, error) {
	log.Debug().Msg("Getting account")

	account, err := a.AccountRepository.Find(accountID)
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}
