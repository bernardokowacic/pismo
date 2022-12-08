package entity

type OperationType struct {
	ID          uint64 `json:"id" gorm:"primaryKey;column:id"`
	Description string `json:"description" binding:"required" gorm:"type:string;size:300;column:content;<-:create;not null"`
}
