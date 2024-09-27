package model

import "gorm.io/gorm"

type Dao struct {
	Orm *gorm.DB
}

/*get grom db*/
func (d *Dao) ORM() *gorm.DB {
	return d.Orm
}
