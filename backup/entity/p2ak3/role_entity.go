package p2ak3

import "gorm.io/gorm"

func (P2ak3Role) TableName() string {
	return "p2ak3_roles"
}

type P2ak3Role struct {
	gorm.Model
	ID   uint64 `gorm:"primaryKey;autoIncrement"`
	Name string
}
