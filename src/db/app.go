package db

import (
	"github.com/go-sql-driver/mysql"

	"ecom/src/config"
)

func GetConfig() mysql.Config {
	return mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
}
