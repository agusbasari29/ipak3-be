package p2ak3

import (
	"time"

	"gorm.io/gorm"
)

func (P2ak3Company) TableName() string {
	return "p2ak3_companies"
}

type P2ak3Company struct {
	gorm.Model
	ID              uint64 `gorm:"primaryKey;autoIncrement"`
	Name            string
	Address         string
	ContactID       uint64
	Contact         P2ak3Contact //`gorm:"foreignKey:ContactID"`
	BusinessID      uint64
	Business        P2ak3Business //`gorm:"foreignKey:BusinessID"`
	RegisteredDate  time.Time
	RegistrationNo  string
	NPWP            string
	EstablishedDate time.Time
}
