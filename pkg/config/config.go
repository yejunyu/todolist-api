package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server       ServerConfig
	Database     DatabaseConfig
	TestDatabase DatabaseConfig `mapstructure:"test_database"`
	JWT          JWTConfig
}
type ServerConfig struct {
	Port int
}
type JWTConfig struct {
	Secret      string `yaml:"secret" mapstructure:"secret"`
	ExpireHours int    `yaml:"expire_hours" mapstructure:"expire_hours"`
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var Cfg *Config

func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// 先读取配置文件
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("failed to read config file: %w", err)
		}
		log.Println("Config file not found, using default values")
	}

	// 启用环境变量自动匹配
	viper.AutomaticEnv()
	// 设置环境变量的前缀，例如 APP_SERVER_PORT
	viper.SetEnvPrefix("APP")
	// 设置环境变量名和键名的替换规则，将 . 替换为 _
	// 这样 viper.Get("database.host") 会自动查找 APP_DATABASE_HOST 环境变量
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err = viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
