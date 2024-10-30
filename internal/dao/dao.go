package dao

import (
	"github.com/strong-defi/strong-defi-backend/internal/model"
	"gorm.io/gorm"
)

func NewDao(orm *gorm.DB) *model.Dao {
	d := &model.Dao{Orm: orm}
	return d
}
