package account

import (
	"pismo/entity"
	"pismo/repository"

	"github.com/rs/zerolog/log"
)

type AccountServiceInterface interface {
	Insert(account entity.Account) (entity.Account, error)
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
