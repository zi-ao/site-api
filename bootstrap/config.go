package bootstrap

import (
	"flag"
	"github.com/zi-ao/site-api/pkg/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
)

// SetupConfig 初始化配置文件
func SetupConfig() *config.Config {
	configPath := *flag.String("config", "./config", "配置文件目录")
	data, err := ioutil.ReadFile(path.Join(configPath, "config.yaml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, config.Global)
	if err != nil {
		panic(err)
	}

	return config.Global
}
