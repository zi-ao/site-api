package bootstrap

import (
	"fmt"
	"github.com/zi-ao/site-api/pkg/config"
	"github.com/zi-ao/site-api/pkg/model"
	"time"
)

func SetupDatabase() {
	db := model.ConnectDB(config.Global.MySQL.DSN())

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// 设置最大连接数
		sqlDB.SetMaxOpenConns(50)
		// 设置最大空闲连接数
		sqlDB.SetMaxIdleConns(10)
		// 设置每个链接的过期时间
		sqlDB.SetConnMaxLifetime(5 * time.Minute)
	}
}
