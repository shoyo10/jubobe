package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	PatientID int       `json:"patient_id" gorm:"column:patient_id;type:int;uniqueIndex;not null"`
	Message   string    `json:"message" gorm:"column:message;type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null"`
}

func (*Order) TableName() string {
	return "orders"
}

type OrderOption struct {
	Filter OrderFilter
}

type OrderFilter struct {
	ID int
}

func (o *OrderFilter) Where(db *gorm.DB) *gorm.DB {
	if o.ID != 0 {
		db = db.Where("id = ?", o.ID)
	}
	return db
}

type UpdateOrderInput struct {
	Message string
}
