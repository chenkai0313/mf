package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	Redis          Redis  `mapstructure:"redis"`
	Server         Server `mapstructure:"server"`
	LogPath        string `mapstructure:"log_path"`
	Debug          bool
	ApiRequestRate ApiRequestRate `mapstructure:"api_request_rate"`
	Consul         Consul         `mapstructure:"consul"`
}

func LoadConfig(debug bool) {
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config error failed to read the configuration file: %s", err))
	}
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unmarshal config err : %s", err.Error()))
	}
	if debug {
		config.Debug = true
		SetMode(DebugMode)
	}
}

type Consul struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Server_Env string

type ApiRequestRate struct {
	DuringTime time.Duration `mapstructure:"during_time"`
	Rate       int64         `mapstructure:"rate"`
}

const (
	Server_Env_Dev  Server_Env = "dev"
	Server_Env_Prod Server_Env = "prod"
)

type Server struct {
	Port string     `mapstructure:"port"`
	Name string     `mapstructure:"name"`
	Host string     `mapstructure:"host"`
	Version string     `mapstructure:"version"`
	Env  Server_Env `mapstructure:"env"`
}

type Redis struct {
	HostName string `mapstructure:"hostname"`
	DB       int    `mapstructure:"database"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

func GetRedis() Redis {
	return config.Redis
}

func GetHttpServer() Server {
	return config.Server
}

func GetLogPath() string {
	return config.LogPath
}

func GetDebug() bool {
	return config.Debug
}

func GetApiRequestRate() ApiRequestRate {
	return config.ApiRequestRate
}

func GetConsul() Consul {
	return config.Consul
}
