package configs

import (
	"github.com/kelseyhightower/envconfig"
)

var osConfig = NewOsConfig()

// OsConfig は環境変数から設定値を取得する際に利用します。
// Configer interfaceを満たす様に実装してください。
type OsConfig struct {
	// config はConfigをメンバーに持ちます。
	config Config
}

// NewOsConfig は環境変数からConfigにセット出来る値を取得してOsConfigのメンバーに追加します。
func NewOsConfig() *OsConfig {
	var c Config
	envconfig.Process("", &c)
	return &OsConfig{config: c}
}

// GetOsConfigInstance はパッケージ内のグローバルスコープに存在するosConfigの値を取得します。
func GetOsConfigInstance() *OsConfig {
	return osConfig
}

// Get はOsConfigのメンバー変数のconfigを取得します。
func (oc *OsConfig) Get() Config {
	return osConfig.config
}
