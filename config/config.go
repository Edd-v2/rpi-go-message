package config

import (
	"os"

	"github.com/spf13/viper"
)

type MongoConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
	Mode string `yaml:"mode"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`  // e.g., "info", "debug", "error"
	Format string `yaml:"format"` // e.g., "json", "text"
}

type AuthConfig struct {
	JWTSecret       string `yaml:"jwt_secret"`
	TokenExpiration int    `yaml:"token_expiration"` // in minutes
}

type MetricsConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"` // e.g., "/metrics"
}

type Config struct {
	Mongo        MongoConfig   `yaml:"mongo"`
	ServerConfig ServerConfig  `yaml:"server"`
	Logging      LoggingConfig `yaml:"logging"`
	Metrics      MetricsConfig `yaml:"metrics"`
	Auth         AuthConfig    `yaml:"auth"`
}

var AppConfig Config

func LoadConfiguration() int {
	//allow override of config path from env
	customPath := os.Getenv("GO_MESSAGE_CONFIG_PATH")
	if customPath != "" {
		viper.AddConfigPath(customPath)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
	}

	viper.SetConfigName("configuration") // configuration.yaml
	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()

	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.mode", "debug")
	viper.SetDefault("metrics.enabled", true)
	viper.SetDefault("metrics.path", "/metrics")

	if err := viper.ReadInConfig(); err != nil {
		return 1
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return 1
	}

	return 0
}
