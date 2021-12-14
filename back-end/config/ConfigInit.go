package config

import (
	"log"

	"github.com/spf13/viper"
)
func InitConfig()(err error) {
	log.Println("初始yaml配置信息。。。")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err=viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return err
	}
	var config Config
	viper.Unmarshal(&config)
	log.Printf("config.Port: %v\n", config.Port)
	log.Printf("config.Mysql: %v\n", config.Mysql)
	log.Println("初始yaml配置完成。。。")
	ConfigFunc=config
	return nil
}
