package config

import "log/slog"

// 共通設定
type CommonConfig struct{}

// ログ設定
type LogConfig struct {
	Options   *slog.HandlerOptions
	IsJsonLog bool
}

// ミドルウェア設定
type MiddlewareConfig struct {
	// `Sec-Fetch-Site` ヘッダーの許可する値
	// なにも設定しない場合すべての値を許可します
	AllowSecFetchSiteValues []string
}

type Config struct {
	Mode string

	LogConfig
	CommonConfig
	MiddlewareConfig
}

type Environment interface {
	GetMode() string
	GetConfig() Config
}

var configs = map[string]func(path string) Environment{
	"local":      SetLocalConfig,
	"production": SetProductionConfig,
}

// mode に応じた設定を取得
// mode が不正な場合はローカル設定を返す
func GetConfig(mode string, path string) Environment {
	setConfigFunc := configs[mode]
	if setConfigFunc == nil {
		return SetLocalConfig(path)
	}

	return setConfigFunc(path)
}
