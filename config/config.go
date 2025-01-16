package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env *viper.Viper
}

func NewConfig() *Config {
	cnf := viper.New()
	cnf.SetConfigFile(".env")
	cnf.AddConfigPath(".")
	cnf.AutomaticEnv()

	log.Println("Using config file", cnf.ConfigFileUsed())
	if err := cnf.ReadInConfig(); err != nil{
		log.Fatal("Can't read file .env file ", err)
	}

	return &Config{Env: cnf}
}