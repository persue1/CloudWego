package dal

import (
	"github.com/persue1/CloudWego/demo/demo_thrift/biz/dal/mysql"
	"github.com/persue1/CloudWego/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
