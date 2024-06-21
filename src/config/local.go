package config

type Local struct {
	Config
}

func (l Local) GetMode() string {
	return l.Mode
}

func (l Local) GetConfig() Config {
	return l.Config
}

func SetLocalConfig() Environment {
	return Local{
		Config: Config{
			Mode: "local",
		},
	}
}
