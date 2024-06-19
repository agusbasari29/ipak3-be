package p2ak3

import "gorm.io/gorm"

type Tabler interface {
	TableName() string
}

func (P2ak3User) TableName() string {
	return "p2ak3_users"
}

type P2ak3User struct {
	gorm.Model
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	Username   string
	Password   string
	FullName   string
	EmployeeID string
	ContactID  uint64
	Contact    P2ak3Contact //`gorm:"foreignKey:ContactID"`
	CompanyID  uint64
	Company    P2ak3Company //`gorm:"foreignKey:CompanyID"`
	RoleID     uint64
	Role       P2ak3Role //`gorm:"references:ID"`
}
