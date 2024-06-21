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

var configs = map[string]func(path string) Environment{
	"local":      SetLocalConfig,
	"production": SetProductionConfig,
}

func GetConfig(mode string, path string) Environment {
	setConfigFunc := configs[mode]
	if setConfigFunc == nil {
		return SetLocalConfig(path)
	}

	return setConfigFunc(path)
}
