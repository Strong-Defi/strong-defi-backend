package model

import "time"

type ScTransactionLog struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	TransType   string    `gorm:"type:varchar(64);default:'';not null;comment:交易类型" json:"trans_type"`
	TransHash   string    `gorm:"type:varchar(128);default:'';not null;comment:交易hash" json:"trans_hash"`
	Gas         int64     `gorm:"default:0;not null;comment:消耗的gas" json:"gas"`
	GasPrice    int64     `gorm:"default:0;not null;comment:gas价格" json:"gas_price"`
	TransStatus int8      `gorm:"default:0;not null;comment:交易状态" json:"trans_status"`
	TransLogs   string    `gorm:"type:varchar(512);default:'';not null;comment:交易日志" json:"trans_logs"`
	TransCost   int64     `gorm:"default:0;not null;comment:总的交易的费用" json:"trans_cost"`
	TransFrom   string    `gorm:"type:varchar(128);default:'';not null;comment:交易from地址" json:"trans_from"`
	TransTo     string    `gorm:"type:varchar(128);default:'';not null;comment:交易to地址" json:"trans_to"`
	Creator     string    `gorm:"type:varchar(64);default:'';not null;comment:创建人" json:"creator"`
	IsDeleted   bool      `gorm:"default:0;not null;comment:是否已删除；1=是，0=否" json:"is_deleted"`
	MTime       time.Time `gorm:"default:CURRENT_TIMESTAMP;not null;comment:更新时间;update:CURRENT_TIMESTAMP" json:"mtime"`
	CTime       time.Time `gorm:"default:CURRENT_TIMESTAMP;not null;comment:创建时间" json:"ctime"`
}

func (ScTransactionLog) TableName() string {
	return "sc_transaction_logs"
}

// CreateScTransLog a new transaction log
func (d *Dao) CreateScTransLog(log *ScTransactionLog) error {
	return d.Orm.Create(log).Error
}

// GetByID Read a transaction log by ID
func (d *Dao) GetByID(id uint64) (*ScTransactionLog, error) {
	var scTransLog ScTransactionLog
	err := d.Orm.First(&scTransLog, id).Error
	return &scTransLog, err
}

// UpdateScTransLog a transaction log
func (d *Dao) UpdateScTransLog(log *ScTransactionLog) error {
	return d.Orm.Save(log).Error
}

// DeleteScTransLogById Delete a transaction log
func (d *Dao) DeleteScTransLogById(id uint64) error {
	return d.Orm.Delete(&ScTransactionLog{}, id).Error
}
