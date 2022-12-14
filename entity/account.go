package entity

type Account struct {
	ID                   uint64  `json:"id" uri:"accountId" gorm:"primaryKey;column:id"`
	DocumentNumber       string  `json:"document_number" gorm:"unique;type:string;size:300;column:document_number;<-:create;not null"`
	AvailableCreditLimit float64 `json:"available_credit_limit" binding:"required" gorm:"type:decimal(10,2);column:available_creadit_limit;not null"`
}
