package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitMysqlConn() (db *gorm.DB) {
	dsn := "root:admin@tcp(116.205.138.224:3306)/Navigation_Web?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	DB = db
	if err != nil {
		fmt.Printf("err:", err)
		log.Printf(err.Error())
		return nil
	}
	return DB
}
