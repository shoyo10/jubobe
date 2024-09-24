package model

import "time"

type Patient struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(32);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`

	Order *Order
}

func (*Patient) TableName() string {
	return "patients"
}
