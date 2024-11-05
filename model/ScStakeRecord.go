package model

import (
	"time"
)

// ScStakeRecord 质押操作记录表
type ScStakeRecord struct {
	ID                uint64    `gorm:"primaryKey;autoIncrement;comment:'主键'" json:"id"`
	StakeType         string    `gorm:"size:64;default:'';not null;comment:'质押类型'" json:"stake_type"`
	PoolID            int64     `gorm:"default:-1;not null;comment:'质押池id'" json:"pool_id"`
	UserWalletAddress string    `gorm:"size:128;default:'';not null;comment:'用户钱包地址'" json:"user_wallet_address"`
	TransHash         string    `gorm:"size:128;default:'';not null;comment:'交易hash'" json:"trans_hash"`
	TransAmount       int64     `gorm:"default:0;not null;comment:'交易金额'" json:"trans_amount"`
	Creator           string    `gorm:"size:64;default:'';not null;comment:'创建人'" json:"creator"`
	IsDeleted         bool      `gorm:"default:false;not null;comment:'是否已删除；1=是，0=否'" json:"is_deleted"`
	MTime             time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null;comment:'更新时间'" json:"mtime"`
	CTime             time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:'创建时间'" json:"ctime"`
}

// TableName 重写表名（如果表名不符合默认规则，可以在这里修改）
func (ScStakeRecord) TableName() string {
	return "staking_transactions"
}

// CreateScStakeRecord 1. 创建一条新的记录
func (d *Dao) CreateScStakeRecord(record ScStakeRecord) error {
	return d.Orm.Create(&record).Error
}
