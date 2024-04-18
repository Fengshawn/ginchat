package models

import (
	"fmt"
	"ginchat/utils"

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

func SearchGroup(userId uint) []GroupBasic {
	groupbasic := make([]GroupBasic, 0)
	objIds := make([]string, 0)
	utils.DB.Where("owner_id = ?", userId).Find(&groupbasic)
	for _, v := range groupbasic {
		fmt.Println(v)
		objIds = append(objIds, v.Name)
	}
	groups := make([]GroupBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&groups)
	return groups
}
