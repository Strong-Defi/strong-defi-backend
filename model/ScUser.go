package model

import (
	"errors"
	"gorm.io/gorm"
	"strong-defi-backend/pkg/log"
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

func (d *Dao) SelectUser(where string, params ...interface{}) (scUser *ScUser, err error) {
	scUser = new(ScUser)
	result := d.ORM().Where(where, params...).First(&scUser)

	//if result.Error != nil {
	//	panic(result.Error)
	//	return nil, result.Error
	//}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return scUser, nil
}

func (d *Dao) SelectUserList(where string, params ...interface{}) ([]ScUser, error) {
	var scUser []ScUser
	result := d.ORM().Where(where, params).Find(&scUser)

	if result.Error != nil {
		panic(result.Error)
	}
	return scUser, nil
}
func (d *Dao) SelectUserByWalletAddress(walletAddress string) (scUser *ScUser, err error) {
	scUser = new(ScUser)
	if err = d.ORM().Where("user_wallet_address = ?", walletAddress).First(&scUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return scUser, nil
}

func (d *Dao) SelectUserPage(pageNum, pageSize int, orderBy, where string, params ...interface{}) (scUsers *[]ScUser, err error) {

	orm := d.ORM().Where(where, params)

	/*默认数据按照ctime倒排*/
	if orderBy != "" {
		orm.Order(orderBy)
	} else {
		orm.Order("mtime desc")
	}

	if pageNum > 0 {
		orm = orm.Limit(pageNum)
	}

	if pageSize > 0 {
		orm = orm.Offset(pageSize)
	}

	if err = orm.Find(&scUsers).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Msgf("scUser.SelectUserPage where(%s) params(%+v) offset(%d) limit(%d) error(%+v)",
				where, params, pageSize, pageNum, err)
			return
		}
	}

	return scUsers, nil
}

func (d *Dao) Create(scUser *ScUser) {
	result := d.ORM().Create(scUser)

	if result.Error != nil {
		panic(result.Error)
	}
}

func (d *Dao) Update(where string, values map[string]interface{}, queryArgs ...interface{}) (int, error) {
	result := d.ORM().Model(&ScUser{}).Where(where, queryArgs).Updates(values)

	if result.Error != nil {
		return 0, result.Error
	} else {
		return int(result.RowsAffected), nil
	}
}

func (d *Dao) SaveScUser(user *ScUser) (err error) {

	if err = d.ORM().Save(user).Error; err != nil {
		log.Error().Msgf("dmDao.SaveContractEntity entity(%+v) error(%+v)", user, err)
	}
	return err
}

func (d *Dao) DeleteScUser(where string, params ...interface{}) error {
	tx := d.ORM().Where(where, params).Delete(&ScUser{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
