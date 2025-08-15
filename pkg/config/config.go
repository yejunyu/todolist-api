package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server       ServerConfig
	Database     DatabaseConfig
	TestDatabase DatabaseConfig `mapstructure:"test_database"`
}
type ServerConfig struct {
	Port int
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

	// --- 关键修改在这里 ---
	// 启用环境变量自动匹配
	viper.AutomaticEnv()
	// 设置环境变量的前缀，例如 APP_SERVER_PORT
	viper.SetEnvPrefix("APP")
	// 设置环境变量名和键名的替换规则，将 . 替换为 _
	// 这样 viper.Get("database.host") 会自动查找 APP_DATABASE_HOST 环境变量
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// --- 修改结束 ---

	if err = viper.ReadInConfig(); err != nil {
		// 让配置文件变为可选，如果找不到也不报错，这样可以完全依赖环境变量
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("failed to read config file: %w", err)
		}
	}

	if err = viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return nil
}
