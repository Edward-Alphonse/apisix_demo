package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var gormDB *gorm.DB

func Init(config *Config) {
	connectDB(config)
}

func connectDB(config *Config) {
	dsn := config.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:         "",
			SingularTable:       true,
			NoLowerCase:         false,
			IdentifierMaxLength: 0,
		},
	})
	if err != nil {
		log.Panic("连接数据库失败", map[string]any{
			"error": err.Error(),
			"dsn":   dsn,
		})
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panic("获取数据库实例失败", map[string]any{
			"error": err.Error(),
			"dsn":   dsn,
		})
	}
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(config.MaxOpenConns) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(config.MaxIdleConns) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于最大空闲数，超过的连接会被连接池关闭。

	gormDB = db
}

func GetDB() *gorm.DB {
	return gormDB
}
