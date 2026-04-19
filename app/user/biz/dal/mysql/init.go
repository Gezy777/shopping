package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/user/model"
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		mysqlHost,
		database,
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