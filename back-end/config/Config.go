package config

type Config struct {
	Server
	Mysql
}

var ConfigFunc Config