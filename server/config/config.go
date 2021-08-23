package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	Server     Server `mapstructure:"server"`
	LogPath    string `mapstructure:"log_path"`
	Consul     Consul `mapstructure:"consul"`
	Debug      bool
	MysqlDbDns string `mapstructure:"mysql_db_dns"`
}

type Consul struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
type Server struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type Db struct {
	Project  string `mapstructure:"project"`
	Driver   string `mapstructure:"driver"`
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Charset  string `mapstructure:"charset"`
	Ssl      string `mapstructure:"ssl"`
}

type Server_Env string

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
	}
}

func GetConsul() Consul {
	return config.Consul
}

func GetServer() Server {
	return config.Server
}

func GetDb() string {
	return config.MysqlDbDns
}

func GetLogPath() string {
	return config.LogPath
}

func GetDebug() bool {
	return config.Debug
}
