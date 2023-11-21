package util

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	PGVersion  string `mapstructure:"PG_VERSION"`
	PGDb       string `mapstructure:"PG_DB"`
	PGUser     string `mapstructure:"PG_USER"`
	PGPassword string `mapstructure:"PG_PASSWORD"`
	PGHost     string `mapstructure:"PG_HOST"`
	PGPort     string `mapstructure:"PG_PORT"`
}

var (
	configSingleton *Config
	configOnce      sync.Once
)

// Initialize config singleton with thread safety
func InitConfig(path, name string) {
	configOnce.Do(func() {
		config, err := loadConfig(path, name)
		if err != nil {
			log.Fatal("cannot load config:", err)
			return
		}
		configSingleton = config
	})
}

// LoadConfig reads configuration from file or env variable
func loadConfig(path, name string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func GetConfigSingleton() *Config {
	return configSingleton
}
