package models

import "github.com/jinzhu/gorm"

type UserBasic struct {
	gorm.Model
	Name string
}
