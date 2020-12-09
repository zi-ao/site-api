package config

import "fmt"

type MySQL struct {
	Host      string `yaml:"host"`
	Port      uint   `yaml:"port"`
	DBName    string `yaml:"dbname"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Charset   string `yaml:"charset"`
	Collation string `yaml:"collation"`
}

func (mysql *MySQL) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&parseTime=True&loc=Local",
		mysql.Username,
		mysql.Password,
		mysql.Host,
		mysql.Port,
		mysql.DBName,
		mysql.Charset,
		mysql.Collation,
	)
}
