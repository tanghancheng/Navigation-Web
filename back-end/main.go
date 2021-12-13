package main

import (
	"Navigation-Web/config"
	"Navigation-Web/dao"
	"Navigation-Web/models"
	"Navigation-Web/routers"
	"fmt"
)

func main() {
	config:=config.InitConfig()
	
	DB := dao.InitMysqlConn(config.Mysql)
	err := DB.AutoMigrate(&models.NavigationInfo{})
	err2 := DB.AutoMigrate(&models.Note{})
	err3 := DB.AutoMigrate(&models.User{})
	if err != nil || err2 != nil || err3 != nil {
		fmt.Println("模型绑定失败  ", err)
	}
	routers.InitRoutes()
}
