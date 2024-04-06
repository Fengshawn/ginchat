package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	OwnerId  uint //哪个用户的关系信息
	TargetId uint // 对应谁
	Type     int  // 对应的类型 0 1 3
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
