package dao

import (
	"Navigation-Web/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysqlConn(mysqlConfig config.Mysql) (db *gorm.DB) {
	fmt.Printf("mysqlConfig: %v\n", mysqlConfig)
	dsn:=fmt.Sprintf("%s:%s@tcp(%s)%s",mysqlConfig.Username,mysqlConfig.Password,mysqlConfig.Addr,mysqlConfig.Database)
	// dsn := "root:admin@tcp(116.205.138.224:3306)/Navigation_Web?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Printf("dsn: %v\n", dsn)
	db, err := gorm.Open(mysql.Open(dsn))
	DB = db
	if err != nil {
		log.Println("err:",err.Error())
		return nil
	}
	return DB
}
