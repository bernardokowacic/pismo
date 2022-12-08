package repository

import (
	"pismo/entity"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type AccountInterface interface {
	Insert(account entity.Account) (entity.Account, error)
}

type accountRepositoryStruct struct {
	DbConn *gorm.DB
}

func NewRepository(dbConn *gorm.DB) AccountInterface {
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
