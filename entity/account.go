package entity

type Account struct {
	ID             uint64 `json:"id" gorm:"primaryKey;column:id"`
	DocumentNumber uint64 `json:"document_number" gorm:"unique;type:uint;column:document_number;<-:create;not null"`
}
