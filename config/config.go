package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config is an exportable type
type Config struct {
	Critical  int `mapstructure:"critical_battery_warning_at"`
	Low       int `mapstructure:"low_battery_warning_at"`
	First     int `mapstructure:"first_custom_warning_at"`
	Second    int `mapstructure:"second_custom_warning_at"`
	Third     int `mapstructure:"third_custom_warning_at"`
	Stability int `mapstructure:"notification_stability"`
}

// GetConfig exports config settings
func GetConfig() *Config {
	viper.SetDefault("critical_battery_warning_at", 10)
	viper.SetDefault("low_battery_warning_at", 30)
	viper.SetDefault("first_custom_warning_at", 50)
	viper.SetDefault("second_custom_warning_at", 75)
	viper.SetDefault("third_custom_warning_at", 90)
	viper.SetDefault("notification_stability", 5)
	viper.SetConfigName("config")                 // name of config file (without extension)
	viper.AddConfigPath("/etc/battery-monitor/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.battery-monitor") // call multiple times to add many search paths
	viper.AddConfigPath(".")                      // optionally look for config in the working directory
	err := viper.ReadInConfig()                   // Find and read the config file
	if err != nil {                               // Handle errors reading the config file
		log.Warn(err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Warn(err)
	}

	return config
}
