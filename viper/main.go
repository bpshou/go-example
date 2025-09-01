package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// 初始化
	ViperInitEnv()
	LoadViperConfig(ViperConfig{
		Name: "config",
		Type: "yaml",
		Path: "./etc/",
		Must: true,
	})
	LoadViperConfig(ViperConfig{
		Name: "user",
		Type: "json",
		Path: "./etc/",
		Must: true,
	})

	json, err := json.Marshal(viper.AllSettings())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))

	user := viper.GetString("USER")
	fmt.Println(user)
}

// 调用AutomaticEnv函数，开启环境变量读取
func ViperInitEnv() {
	viper.AutomaticEnv()
}

type ViperConfig struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
	Must bool   `json:"must"`
}

func LoadViperConfig(config ViperConfig) error {
	viper.SetConfigName(config.Name)
	viper.SetConfigType(config.Type)
	viper.AddConfigPath(config.Path)
	err := viper.MergeInConfig()
	if err != nil {
		if config.Must {
			panic(err)
		}
		return err
	}
	return nil
}
