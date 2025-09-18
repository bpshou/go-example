package core

import "gin_app/app/core/config"

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
}
