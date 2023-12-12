package config

import (
	"gamelib/internal/storage/db"
	"gamelib/pkg/web"
	"github.com/spf13/viper"
)

type Config struct {
	Server  web.ServerConfig
	Storage db.StorageConfig
}

func ParseConfig(env string) error {
	viper.AddConfigPath("configs/" + env)
	viper.SetConfigName("common")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
