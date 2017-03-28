package volume

import "github.com/spf13/viper"

const (
	vMax   = "volume.max"
	vMin   = "volume.min"
	vCard  = "volume.card"
	vMixer = "volume.mixer"
)

type Configurer interface {
	Max() float64
	Min() float64
	Card() string
	Mixer() string
}

func init() {
	// Set Defaults
	viper.SetDefault(vMax, 50)
	viper.SetDefault(vMin, 10)
	viper.SetDefault(vCard, "default")
	viper.SetDefault(vMixer, "PCM")
	// Bind Environment Variables
	viper.BindEnv(
		vMax,
		vMin,
		vCard,
		vMixer)
}

type Config struct{}

func (c Config) Max() float64 {
	return viper.GetFloat64(vMax)
}

func (c Config) Min() float64 {
	return viper.GetFloat64(vMin)
}

func (c Config) Card() string {
	return viper.GetString(vCard)
}

func (c Config) Mixer() string {
	return viper.GetString(vMixer)
}
