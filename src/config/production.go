package config

type Production struct {
	Config
}

func (l Production) GetMode() string {
	return l.Mode
}

func (l Production) GetConfig() Config {
	return l.Config
}

func SetProductionConfig() Environment {
	return Local{
		Config: Config{
			Mode: "production",

			CommonConfig: GetCommonConfig(),
		},
	}
}
