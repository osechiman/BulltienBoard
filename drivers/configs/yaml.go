package configs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var yamlConfig = NewYamlConfig()

const configFilePath = "./drivers/configs/config.yaml"

// YamlConfig はyamlファイルから設定値を取得する際に利用します。
// Configer interfaceを満たす様に実装してください。
type YamlConfig struct {
	config Config
}

// NewYamlConfig はyamlファイルからConfigにセット出来る値を取得してYamlConfigのメンバーに追加します。
func NewYamlConfig() *YamlConfig {
	var c Config
	body, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = yaml.Unmarshal(body, &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &YamlConfig{config: c}
}

// GetYamlConfigInstance はパッケージ内のグローバルスコープに存在するyamlConfigの値を取得します。
func GetYamlConfigInstance() *YamlConfig {
	return yamlConfig
}

// Get はyamlConfigのメンバー変数のconfigを取得します。
func (oc *YamlConfig) Get() Config {
	return yamlConfig.config
}
