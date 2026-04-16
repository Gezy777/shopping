package dal

import (
	"cwgo_try/biz/dal/mysql"
	"cwgo_try/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
