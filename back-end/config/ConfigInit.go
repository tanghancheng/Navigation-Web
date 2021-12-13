package config

import (
	"log"

	"github.com/spf13/viper"
)
func InitConfig()(config *Config) {
	log.Println("初始yaml配置信息。。。")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err:=viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	
	viper.Unmarshal(&config)
	log.Printf("config.Port: %v\n", config.Port)
	log.Printf("config.Mysql: %v\n", config.Mysql)
	log.Println("初始yaml配置完成。。。")
	return config
}
