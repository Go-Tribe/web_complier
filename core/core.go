package core

import (
	"sync"
	c "web_complier/configs"
)

var once sync.Once

func InitCore() {
	once.Do(func() {
		if c.Config.Mysql.Enable == true {
			// 初始化 mysql
			initMysql()
		}

		if c.Config.Redis.Enable == true {
			// 初始化 redis
			initRedis()
		}
		// 初始化日志
		InitLogger()
	})
}
