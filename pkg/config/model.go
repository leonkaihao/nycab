package config

import "github.com/spf13/viper"

// Service is what service.nycab need
type Service struct {
	DBHost   string `json:"dbHost"`
	DBPort   string `json:"dbPort"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// API is what api.nycab need
type API struct {
	SvcURL   string `json:"svcUrl"`
	RedisURL string `json:"redisUrl"`
}

// Cli is waht cli.nycab need
type Cli struct {
	APIHost string `json:"apiHost"`
	APIPort string `json:"apiPort"`
}

// GetServiceConfig ...
func GetServiceConfig(v *viper.Viper) (*Service, error) {
	c := &Service{}
	err := v.UnmarshalKey("service", c)
	return c, err
}

// GetAPIConfig ...
func GetAPIConfig(v *viper.Viper) (*API, error) {
	c := &API{}
	err := v.UnmarshalKey("api", c)
	return c, err
}

// GetCliConfig ...
func GetCliConfig(v *viper.Viper) (*Cli, error) {
	c := &Cli{}
	err := v.UnmarshalKey("cli", c)
	return c, err
}
