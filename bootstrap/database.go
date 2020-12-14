package bootstrap

import (
	"fmt"
	"github.com/zi-ao/site-api/app/models"
	"github.com/zi-ao/site-api/pkg/config"
	"github.com/zi-ao/site-api/pkg/model"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

func SetupDatabase() {
	var mode gormLogger.LogLevel
	if config.Global.Debug {
		mode = gormLogger.Info
	} else {
		mode = gormLogger.Error
	}
	gConf := &gorm.Config{
		Logger: gormLogger.Default.LogMode(mode),
	}
	db := model.ConnectDB(config.Global.MySQL.DSN(), gConf)

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

	migrate(db)
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Category{},
		&models.Tag{},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
}
