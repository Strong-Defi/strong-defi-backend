package dao

type Db struct {
	Dsn         string // 是否输出到控制台
	Active      int    //链接池数量
	Idle        int    //存货
	IdleTimeout int    //超时时间
}
