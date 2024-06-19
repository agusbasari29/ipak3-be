package ipak3

import (
	"time"

	"gorm.io/gorm"
)

func (Ipak3Company) TableName() string {
	return "ipak3_companies"
}

type Ipak3Company struct {
	gorm.Model
	ID              uint64 `gorm:"primaryKey;autoIncrement"`
	Name            string
	Address         string
	ContactID       uint64
	Contact         Ipak3Contact //`gorm:"foreignKey:ContactID"`
	BusinessID      uint64
	Business        Ipak3Business //`gorm:"foreignKey:BusinessID"`
	RegisteredDate  time.Time
	RegistrationNo  string
	NPWP            string
	EstablishedDate time.Time
}
