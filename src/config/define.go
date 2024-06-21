package config

type CommonConfig struct {
}

type Config struct {
	Mode string

	CommonConfig
}

type Environment interface {
	GetMode() string
	GetConfig() Config
}
