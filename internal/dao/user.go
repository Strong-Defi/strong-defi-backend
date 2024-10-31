package dao

import (
	"context"
	"errors"
	"github.com/strong-defi/strong-defi-backend/internal/model"
	"gorm.io/gorm"
)

func (d *Dao) SelectUser(where string, params ...interface{}) (scUser *model.ScUser, err error) {
	scUser = new(model.ScUser)

	result := d.db.Where(where, params).First(&scUser)

	if result.Error != nil {
		panic(result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return scUser, nil
}

func (d *Dao) SelectUserList(where string, params ...interface{}) ([]*model.ScUser, error) {
	var scUser []*model.ScUser
	result := d.db.Where(where, params).Find(&scUser)

	if result.Error != nil {
		panic(result.Error)
	}
	return scUser, nil
}
func (d *Dao) SelectUserByWalletAddress(walletAddress string) (scUser *model.ScUser, err error) {
	scUser = new(model.ScUser)
	if err = d.db.Where("user_wallet_address = ?", walletAddress).First(&scUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return scUser, nil
}

func (d *Dao) SelectUserPage(pageNum, pageSize int, orderBy, where string, params ...interface{}) (scUsers *[]*model.ScUser, err error) {

	orm := d.db.Where(where, params)

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
			log.Error("scUser.SelectUserPage where(%s) params(%+v) offset(%d) limit(%d) error(%+v)",
				where, params, pageSize, pageNum, err)
			return
		}
	}

	return scUsers, nil
}

func (d *Dao) Create(scUser *model.ScUser) {
	result := d.db.Create(scUser)

	if result.Error != nil {
		panic(result.Error)
	}
}

func (d *Dao) Update(where string, values map[string]interface{}, queryArgs ...interface{}) (int, error) {
	result := d.db.Model(&model.ScUser{}).Where(where, queryArgs).Updates(values)

	if result.Error != nil {
		return 0, result.Error
	} else {
		return int(result.RowsAffected), nil
	}
}

func (d *Dao) SaveScUser(user *model.ScUser) (err error) {

	if err = d.db.Save(user).Error; err != nil {
		log.Error("dmDao.SaveContractEntity entity(%+v) error(%+v)", user, err)
	}
	return err
}

func (d *Dao) DeleteScUser(where string, params ...interface{}) error {
	tx := d.db.Where(where, params).Delete(&model.ScUser{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d *Dao) FirstOrCreateUser(ctx context.Context, user *model.ScUser) (fUser *model.ScUser, err error) {
	fUser = new(model.ScUser)
	err = d.db.WithContext(ctx).FirstOrCreate(user, user).Error
	if err != nil {
		log.Error("[FirstOrCreateUser] FirstOrCreate err:(%v)", err)
		return fUser, err
	}
	return fUser, nil
}
