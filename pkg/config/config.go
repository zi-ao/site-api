package config

// Global 全局配置
var Global = &Config{}

// Config 系统配置
type Config struct {
	Name  string `yaml:"name"`
	Debug bool   `yaml:"debug"`
	Port  uint   `yaml:"port"`
	URL   string `yaml:"url"`
	Key   string `yaml:"key"`

	LogPath string `yaml:"log_path"`

	MySQL MySQL `yaml:"mysql,flow"`
}
