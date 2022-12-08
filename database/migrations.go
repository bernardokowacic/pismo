package database

import (
	"pismo/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Migrator().CreateTable(&entity.Account{})
	db.Migrator().CreateTable(&entity.OperationType{})
	db.Migrator().CreateTable(&entity.Transaction{})
}
