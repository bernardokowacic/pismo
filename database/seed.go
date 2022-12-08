package database

import (
	"pismo/entity"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	operationTypes := []entity.OperationType{
		{
			Description: "Compra a vista",
		},
		{
			Description: "Compra parcelada",
		},
		{
			Description: "Saque",
		},
		{
			Description: "Pagamento",
		},
	}

	db.Create(&operationTypes)
}
