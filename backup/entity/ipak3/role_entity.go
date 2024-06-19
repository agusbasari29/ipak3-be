package ipak3

import "gorm.io/gorm"

func (Ipak3Role) TableName() string {
	return "ipak3_roles"
}

type Ipak3Role struct {
	gorm.Model
	ID   uint64 `gorm:"primaryKey;autoIncrement"`
	Name string
}
