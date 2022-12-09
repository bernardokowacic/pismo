package repository

import (
	"pismo/entity"

	"gorm.io/gorm"
)

type OperationTypeInterface interface {
	Find(operationTypeID uint64) (entity.OperationType, error)
}

type operationTypeRepositoryStruct struct {
	DbConn *gorm.DB
}

func NewOperationTypeRepository(dbConn *gorm.DB) OperationTypeInterface {
	return &operationTypeRepositoryStruct{DbConn: dbConn}
}

func (a *operationTypeRepositoryStruct) Find(operationTypeID uint64) (entity.OperationType, error) {
	var operationTypeSearch entity.OperationType

	err := a.DbConn.Where("id = ?", operationTypeID).First(&operationTypeSearch).Error
	if err != nil {
		return operationTypeSearch, err
	}

	return operationTypeSearch, nil
}
