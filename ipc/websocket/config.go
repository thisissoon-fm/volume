package websocket

import (
	"time"

	"github.com/spf13/viper"
)

const (
	vHost     = "websocket.host"
	vScheme   = "websocket.scheme"
	vPath     = "websocket.path"
	vRetry    = "websocket.retry"
	vUsername = "websocket.username"
	vPassword = "websocket.password"
)

func init() {
	viper.SetDefault(vHost, "localhost:8000")
	viper.SetDefault(vScheme, "ws")
	viper.SetDefault(vPath, "/")
	viper.SetDefault(vRetry, "5s")
}

type Configurer interface {
	Host() string
	Scheme() string
	Path() string
	Retry() time.Duration
	Username() string
	Password() string
}

type Config struct{}

func (c Config) Host() string {
	viper.BindEnv(vHost)
	return viper.GetString(vHost)
}

func (c Config) Scheme() string {
	viper.BindEnv(vScheme)
	return viper.GetString(vScheme)
}

func (c Config) Path() string {
	viper.BindEnv(vPath)
	return viper.GetString(vPath)
}

func (c Config) Retry() time.Duration {
	viper.BindEnv(vRetry)
	return viper.GetDuration(vRetry)
}

func (c Config) Username() string {
	viper.BindEnv(vUsername)
	return viper.GetString(vUsername)
}

func (c Config) Password() string {
	viper.BindEnv(vPassword)
	return viper.GetString(vPassword)
}

func NewConfig() Config {
	return Config{}
}
