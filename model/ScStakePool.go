package model

import (
	"errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"time"
)

type StakePool struct {
	ID                   uint64    `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	PoolID               int64     `gorm:"default:-1;not null;comment:质押池id" json:"pool_id"`
	PoolCode             string    `gorm:"type:varchar(128);default:'';not null;comment:质押池唯一code" json:"pool_code"`
	TokenAddress         string    `gorm:"type:varchar(128);default:'';not null;comment:token地址" json:"token_address"`
	LockStakeBlockNumber int64     `gorm:"default:0;not null;comment:解锁质押的区块数" json:"lock_stake_block_number"`
	MinStakeAmount       uint64    `gorm:"type:bigint(64);default:0;not null;comment:最小质押金额,单位：wei" json:"min_stake_amount"`
	Weight               int64     `gorm:"default:0;not null;comment:池权重" json:"weight"`
	Creator              string    `gorm:"type:varchar(64);default:'';not null;comment:创建人" json:"creator"`
	IsDeleted            bool      `gorm:"default:0;not null;comment:是否已删除；1=是，0=否" json:"is_deleted"`
	MTime                time.Time `gorm:"default:CURRENT_TIMESTAMP;not null;comment:更新时间;update:CURRENT_TIMESTAMP" json:"mtime"`
	CTime                time.Time `gorm:"default:CURRENT_TIMESTAMP;not null;comment:创建时间" json:"ctime"`
}

func (StakePool) TableName() string {
	return "sc_stake_pool"
}

// 创建质押池
func (d *Dao) CreateStakePool(stakePool *StakePool) error {
	return d.Orm.Create(stakePool).Error
}

// 获取单个质押池
func (d *Dao) GetStakePoolByID(id uint64) (stakePool *StakePool, err error) {
	if err = d.Orm.First(&stakePool, id).Error; err != nil {
		return nil, err
	}
	return stakePool, nil
}

func (d *Dao) GetStakePoolByCode(poolCode string) (stakePool *StakePool, err error) {
	where := "pool_code=? and is_deleted=0"
	if err = d.Orm.Where(where, poolCode).First(&stakePool).Error; err != nil {
		return nil, err
	}
	return stakePool, nil
}

// 获取所有质押池
func (d *Dao) SelectStakePoolPage(pageSize, pageNum int, orderBy, where string, params ...interface{}) (stakePools *[]StakePool, err error) {

	tx := d.ORM().Where(where, params)
	if orderBy != "" {
		tx.Order(orderBy)
	} else {
		tx.Order("mtime desc")
	}

	if pageNum > 0 {
		tx.Limit(pageNum)
	}
	if pageSize > 0 {
		tx.Offset(pageSize)
	}

	if err = tx.Find(&stakePools).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Msgf("stakePool.SelectPoolPage where(%s) params(%+v) offset(%d) limit(%d) error(%+v)",
				where, params, pageSize, pageNum, err)
			return
		}
	}
	return stakePools, nil
}

// 更新质押池
func (d *Dao) UpdateStakePool(id uint64, updates map[string]interface{}) error {
	return d.Orm.Model(&StakePool{}).Where("id = ?", id).Updates(updates).Error
}

// 删除质押池
func (d *Dao) DeleteStakePool(id uint64) error {
	return d.Orm.Delete(&StakePool{}, id).Error
}
