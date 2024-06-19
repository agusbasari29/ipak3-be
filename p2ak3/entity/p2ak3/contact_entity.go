package p2ak3

import "gorm.io/gorm"

func (P2ak3Contact) TableName() string {
	return "p2ak3_contacts"
}

type P2ak3Contact struct {
	gorm.Model
	ID    uint64 `gorm:"primaryKey;autoIncrement"`
	Email string
	Phone string
	Fax   string
}
