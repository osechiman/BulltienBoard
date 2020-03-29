package configs

// Config は設定値を管理するためのStructです。
type Config struct {
	// Environment は実行環境を判別するために使います。 ex.) production or staging or develop
	Environment string `yaml:"environment" envconfig:"ENVIRONMENT" default:"develop"`
}
