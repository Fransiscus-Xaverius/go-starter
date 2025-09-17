package factory

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MakeGormDBConnection() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"), viper.GetString("mysql.password"), viper.GetString("mysql.host"), viper.GetInt("mysql.port"), viper.GetString("mysql.dbname"),
	)

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}

	// Get underlying sql.DB to set connection pool params
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// Set connection pool settings
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxidleconn"))                                   // number of idle connections
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxopenconn"))                                   // maximum open connections
	sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt64("mysql.maxlifetime")) * time.Minute) // connections older than this are closed

	return db
}
