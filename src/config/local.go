package config

import "log/slog"

type Local struct {
	Config
}

func (l Local) GetMode() string {
	return l.Mode
}

func (l Local) GetConfig() Config {
	return l.Config
}

func SetLocalConfig(path string) Environment {
	return Local{
		Config: Config{
			Mode: "local",

			LogConfig: LogConfig{
				Options: &slog.HandlerOptions{
					Level: slog.LevelDebug,
				},
			},
			CommonConfig: GetCommonConfig(),
		},
	}
}
