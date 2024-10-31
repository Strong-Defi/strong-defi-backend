package model

import (
	"time"
)

type ScUser struct {
	ID                uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserWalletAddress string    `gorm:"type:varchar(128);not null" json:"user_wallet_address"`
	UserUID           string    `gorm:"type:varchar(128);not null" json:"user_uid"`
	UserTelephone     string    `gorm:"type:varchar(32);not null" json:"user_telephone"`
	UserName          string    `gorm:"type:varchar(64);not null" json:"user_name"`
	UserPassword      string    `gorm:"type:varchar(64);not null" json:"user_password"`
	UserEmail         string    `gorm:"type:varchar(64);not null" json:"user_email"`
	UserCountry       string    `gorm:"type:varchar(128);not null" json:"user_country"`
	UserAddress       string    `gorm:"type:varchar(128);not null" json:"user_address"`
	Remark            string    `gorm:"type:varchar(512);not null" json:"remark"`
	IsDeleted         int       `gorm:"type:tinyint(1);default:0;not null" json:"is_deleted"`
	MTime             time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;not null;autoUpdateTime" json:"mtime"`
	CTime             time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;not null;autoCreateTime" json:"ctime"`
}

func (ScUser) TableName() string {
	return "sc_user"
}
