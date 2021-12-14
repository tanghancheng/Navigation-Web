package main

import (
	"Navigation-Web/config"
	"Navigation-Web/dao"
	"Navigation-Web/models"
	"Navigation-Web/routers"
	"fmt"
	"log"
)

func main() {
	configErr:=config.InitConfig()
	if configErr != nil {
		log.Println("configErr 初始配置文件失败",configErr)
		return
	}
	DB := dao.InitMysqlConn()
	err := DB.AutoMigrate(&models.NavigationInfo{})
	err2 := DB.AutoMigrate(&models.Note{})
	err3 := DB.AutoMigrate(&models.User{})
	if err != nil || err2 != nil || err3 != nil {
		fmt.Println("模型绑定失败  ", err)
		return
	}
	routers.InitRoutes()
}
