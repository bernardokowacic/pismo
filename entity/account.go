package entity

type Account struct {
	ID             uint64 `json:"id" gorm:"primaryKey;column:id"`
	DocumentNumber uint64 `json:"user_id" gorm:"type:uint;column:document_number;<-:create;not null"`
}
