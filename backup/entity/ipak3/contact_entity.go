package ipak3

import "gorm.io/gorm"

func (Ipak3Contact) TableName() string {
	return "ipak3_contacts"
}

type Ipak3Contact struct {
	gorm.Model
	ID    uint64 `gorm:"primaryKey;autoIncrement"`
	Email string
	Phone string
	Fax   string
}
