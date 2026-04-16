package mysql

import (
	"cwgo_test/biz/model"
	"cwgo_test/conf"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {

	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
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
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{})
	fmt.Printf("%#v", DB.Debug().Exec("select version()"))

}
