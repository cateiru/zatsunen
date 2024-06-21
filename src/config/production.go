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
					AddSource: true,
					Level:     slog.LevelInfo,
					// Google Cloud でのログ出力に必要な設定
					ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
						switch a.Key {
						case slog.MessageKey:
							a = slog.Attr{
								Key:   "message",
								Value: a.Value,
							}
						case slog.LevelKey:
							a = slog.Attr{
								Key:   "severity",
								Value: a.Value,
							}
						case slog.SourceKey:
							a = slog.Attr{
								Key:   "logging.googleapis.com/sourceLocation",
								Value: a.Value,
							}
						}
						return a
					},
				},
			},
			CommonConfig: GetCommonConfig(),
		},
	}
}
