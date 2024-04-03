package config

var config Config

type Config struct {
	Logger Logger `long:"logger" yaml:"logger,omitempty" json:"logger"`
}

type Logger struct {
	Level string `long:"level" yaml:"level,omitempty" json:"level,omitempty"`
}

func GetLogLevel() string {
	return GetConfig().Logger.Level
}

func GetConfig() Config {
	config.Logger.Level = "debug"
	return config
}
