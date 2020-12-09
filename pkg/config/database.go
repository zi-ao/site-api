package config

type MySQL struct {
	Host      string `yaml:"host"`
	Port      uint   `yaml:"port"`
	DBName    string `yaml:"dbname"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Charset   string `yaml:"charset"`
	Collation string `yaml:"collation"`
}
