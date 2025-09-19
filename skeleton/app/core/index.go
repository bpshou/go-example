package core

import (
	"gin_app/app/core/config"
	"gin_app/app/core/db"
	"gin_app/app/core/rds"
	"log"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Init() {
	config.ViperInitEnv()
	config.LoadViperConfig(config.ViperConfig{
		Name: "config",
		Type: "yaml",
		Path: "./etc",
		Must: true,
	})
	config.LoadViperConfig(config.ViperConfig{
		Name: "log",
		Type: "yaml",
		Path: "./etc",
		Must: true,
	})
	config.LoadViperConfig(config.ViperConfig{
		Name: "jwt",
		Type: "yaml",
		Path: "./etc",
		Must: false,
	})
	config.LoadViperConfig(config.ViperConfig{
		Name: "mysql",
		Type: "yaml",
		Path: "./etc",
		Must: false,
	})
	config.LoadViperConfig(config.ViperConfig{
		Name: "redis",
		Type: "yaml",
		Path: "./etc",
		Must: false,
	})

	_, err := db.NewMySQL(viper.GetString("mysql.dsn"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}

	rds.NewRedis(viper.GetString("redis.addr"), viper.GetString("redis.password"), viper.GetInt("redis.db"))
}
