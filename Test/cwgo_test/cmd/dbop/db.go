package main

import (
	"cwgo_test/biz/dal"
	"cwgo_test/biz/dal/mysql"
	"cwgo_test/biz/model"
	//"fmt"

	"github.com/joho/godotenv"
)	

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	//create data
	mysql.DB.Create(&model.User{Email: "1768791459@qq.com", Password: "123456"})
	//update
	//mysql.DB.Model(&model.User{}).Where("email = ?", "1768791459@qq.com").Update("password", "222222")
	//read
	//var row model.User
	//mysql.DB.Model(&model.User{}).Where("email = ?", "1768791459@qq.com").First(&row)
	//fmt.Printf("row: %+v\n", row)
	//delete
	//mysql.DB.Unscoped().Where("email = ?", "1768791459@qq.com").Delete(&model.User{})
}