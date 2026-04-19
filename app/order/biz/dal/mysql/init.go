package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	"github.com/cloudwego/biz-demo/gomall/app/order/conf"
	"github.com/cloudwego/kitex/pkg/klog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)
func Init() {
	mysqlHost := os.Getenv("MYSQL_IPADDRESS")
	if mysqlHost == "" {
		mysqlHost = "localhost"
	}
	username := os.Getenv("MYSQL_USERNAME")
	if username == "" {
		username = "gomall"
	}
	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		password = "gomall123"
	}
	database := os.Getenv("MYSQL_DATABASE")
	if database == "" {
		database = "gomall"
	}
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, username, password, mysqlHost, database)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		if err := DB.AutoMigrate(&model.Order{}, &model.OrderItem{}); err != nil {
			klog.Error(err)
		}
	}
}