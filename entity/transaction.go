package entity

import (
	"time"
)

type Transaction struct {
	ID              uint64    `json:"id" gorm:"primaryKey;column:id"`
	AccountID       uint64    `json:"account_id" binding:"required" gorm:"type:uint;column:account_id;<-:create;not null"`
	OperationTypeID uint64    `json:"operation_type_id" binding:"required" gorm:"type:uint;column:operation_type_id;<-:create;not null"`
	Amount          float64   `json:"amount" binding:"required" gorm:"type:decimal(10,2);column:amount;not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt       time.Time `json:"-" gorm:"autoUpdateTime;not null"`
}
