package ipak3

import "gorm.io/gorm"

type Tabler interface {
	TableName() string
}

func (Ipak3User) TableName() string {
	return "ipak3_users"
}

type Ipak3User struct {
	gorm.Model
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	Username   string
	Password   string
	FullName   string
	EmployeeID string
	ContactID  uint64
	Contact    Ipak3Contact //`gorm:"foreignKey:ContactID"`
	CompanyID  uint64
	Company    Ipak3Company //`gorm:"foreignKey:CompanyID"`
	RoleID     uint64
	Role       Ipak3Role //`gorm:"references:ID"`
}
