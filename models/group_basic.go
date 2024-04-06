package models

import (
	"gorm.io/gorm"
)

// 群信息
type GroupBasic struct {
	gorm.Model
	Name    string //群名
	OwnerId uint
	Icon    string
	Desc    string
}

func (table *GroupBasic) TableName() string {
	return "groupbasic"
}
