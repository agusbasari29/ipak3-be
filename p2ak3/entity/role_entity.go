package entity

import "gorm.io/gorm"

func (Role) TableName() string {
	return "p2ak3_roles"
}

type Role struct {
	gorm.Model
	ID   uint64 `gorm:"primaryKey;autoIncrement"`
	Name string
}
