// Handles Reading in Configuration Files

package config

import (
	"os"
	"strings"
	"volume/log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Configuration defaults
func init() {
	viper.SetTypeByDefaultValue(true)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/sfm/volume")
	viper.AddConfigPath("$HOME/.config/sfm/volume")
	viper.SetEnvPrefix("SFM_VOLUME")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

// Bind CLI Flag for config path override
func BindConfigPathFlag(flag *pflag.Flag) {
	viper.BindPFlag("config.path", flag)
}

// Read in configuration from file
func ReadInConfig() {
	// Absolute configuration file
	if _, err := os.Stat(viper.GetString("config.path")); err == nil {
		viper.SetConfigFile(viper.GetString("config.path"))
	}
	// Read in configuration
	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Warn("error loading configuration")
	}
}
