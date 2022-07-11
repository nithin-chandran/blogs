package config

var _conf *Config

// Config reprsents all the configs for the project.
type Config struct {
	Port string `json:"port"`
}

// InitConfig initializes the config.
func InitConfig() {
	_conf = new(Config)
	_conf.Port = "3000"
}

// GetConfig gets the config.
func GetConfig() *Config {
	return _conf
}
