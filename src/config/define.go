package config

type Config struct {
	Mode string
}

type Environment interface {
	GetMode() string
	GetConfig() Config
}
