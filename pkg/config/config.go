package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// LoadConfig for service based on env
// Local: load a local yaml
// Dev or Prod: env variables
func LoadConfig(svc string) (v *viper.Viper, err error) {
	v = viper.New()
	v.SetConfigFile(fmt.Sprintf("config/%s.yaml", svc))
	err = v.ReadInConfig()
	return
}

// GetLocalConfig gets viper
func GetLocalConfig(path string) (v *viper.Viper, err error) {
	v = viper.New()
	v.SetConfigFile(path)
	err = v.ReadInConfig()
	return
}
