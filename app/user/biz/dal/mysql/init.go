package mysql

import (
	"fmt"
	"os"

	//"github.com/cloudwego/biz-demo/gomall/app/user/conf"
	"github.com/cloudwego/biz-demo/gomall/app/user/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"/*conf.GetConf().MySQL.DSN*/,
	os.Getenv("MYSQL_USERNAME"),
	os.Getenv("MYSQL_PASSWORD"),
	os.Getenv("MYSQL_IPADDRESS"),
	os.Getenv("MYSQL_DATABASE"),
)

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
