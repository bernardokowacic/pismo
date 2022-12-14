package repository

import (
	"pismo/entity"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type AccountInterface interface {
	Insert(account entity.Account) (entity.Account, error)
	Find(accountID uint64) (entity.Account, error)
	UpdateBalance(accountID uint64, newBalance float64) error
}

type accountRepositoryStruct struct {
	DbConn *gorm.DB
}

func NewAccountRepository(dbConn *gorm.DB) AccountInterface {
	return &accountRepositoryStruct{DbConn: dbConn}
}

func (a *accountRepositoryStruct) Insert(account entity.Account) (entity.Account, error) {
	result := a.DbConn.Model(&entity.Account{}).Create(&account)
	if result.Error != nil {
		log.Error().Msg(result.Error.Error())
		return entity.Account{}, result.Error
	}

	return account, nil
}

func (a *accountRepositoryStruct) Find(accountID uint64) (entity.Account, error) {
	var accountSearch entity.Account

	err := a.DbConn.Where("id = ?", accountID).First(&accountSearch).Error
	if err != nil {
		log.Error().Msg(err.Error())
		return accountSearch, err
	}

	return accountSearch, nil
}

func (a *accountRepositoryStruct) UpdateBalance(accountID uint64, newBalance float64) error {
	err := a.DbConn.Model(&entity.Account{}).Where("id = ?", accountID).Update("available_creadit_limit", newBalance).Error
	if err != nil {
		return err
	}

	return nil
}
