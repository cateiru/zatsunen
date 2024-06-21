package config

import "log/slog"

type Production struct {
	Config
}

func (l Production) GetMode() string {
	return l.Mode
}

func (l Production) GetConfig() Config {
	return l.Config
}

func SetProductionConfig(path string) Environment {
	return Local{
		Config: Config{
			Mode: "production",

			LogConfig: LogConfig{
				Options: &slog.HandlerOptions{
					Level: slog.LevelError,
				},
			},
			CommonConfig: GetCommonConfig(),
		},
	}
}
