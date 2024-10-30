package dao

import "gorm.io/gorm"

type Db struct {
	Dsn         string // 是否输出到控制台
	Active      int    //链接池数量
	Idle        int    //存货
	IdleTimeout int    //超时时间
}

func NewDB() *gorm.DB {
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//if err != nil {
	//	logs.Error("failed to connect database")
	//	return nil
	//}
	//return db
	return nil
}
