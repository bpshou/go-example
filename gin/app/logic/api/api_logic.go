package api

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ApiLogic() gin.H {

	host := viper.Get("Database.Host")
	port := viper.Get("Database.Port")
	username := viper.Get("Database.Username")
	password := viper.Get("Database.Password")
	name := viper.Get("Database.Name")

	slog.Info("api logic", "host", host)
	slog.Info("api logic", "port", port)
	slog.Info("api logic", "username", username)
	slog.Info("api logic", "password", password)
	slog.Info("api logic", "name", name)

	return gin.H{
		"message":  "Hello, World!",
		"host":     host,
		"port":     port,
		"username": username,
		"password": password,
		"name":     name,
	}
}
