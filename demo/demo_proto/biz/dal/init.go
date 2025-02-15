package dal

import (
	"github.com/persue1/CloudWego/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
