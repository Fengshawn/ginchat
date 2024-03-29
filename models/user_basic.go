package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time `gorm:"default:null"`
	HeartbeatTime time.Time `gorm:"default:null"`
	LoginOutTime  time.Time `gorm:"column:login_out_time;default:null" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
