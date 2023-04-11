package configs

import (
	_ "embed"
	"fmt"
	"github.com/spf13/viper"
	"runtime"
)

type Config struct {
	App Logger `mapstructure:",squash"`
}

var Cfg *Config

func init() {
	Cfg = LoadConfig()
}

func LoadConfig() *Config {
	// Initialize viper
	cfg := &Config{}

	v := viper.New()

	_, filePath, _, _ := runtime.Caller(0)
	configFile := filePath[:len(filePath)-9]

	v.SetConfigFile(configFile + "config" + ".yaml")

	// Load the config file
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	// Unmarshal the config into a struct
	if err := v.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %s", err))
	}

	// Use the config values
	return cfg
}
