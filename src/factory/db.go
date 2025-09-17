package factory

import (
	"fmt"
	"time"

	"github.com/cde/go-example/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MakeGormDBConnection(cfg *config.Config) *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQLUsername, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDbName,
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
	sqlDB.SetMaxIdleConns(cfg.MySQLMaxIdleConnection)                               // number of idle connections
	sqlDB.SetMaxOpenConns(cfg.MySQLMaxOpenConnection)                               // maximum open connections
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MySQLConnMaxLifetime) * time.Minute) // connections older than this are closed

	return db
}
