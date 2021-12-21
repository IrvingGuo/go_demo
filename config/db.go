package config

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db *gorm.DB
)

type mysqlConf struct {
	DSN          string `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

// TODO: create database if not exists
func initMysql() {
	Logger.Info("Connecting MySQL...")
	var err error
	mysqlConfig := mysql.Config{
		DSN:                       Conf.Mysql.DSN, // DSN data source name
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,  // Not support before MySQL 5.6
		DontSupportRenameIndex:    true,  // Not support for MariaDB and before MySQL 5.7
		DontSupportRenameColumn:   true,  // Not support for MariaDB and before MySQL 8
		SkipInitializeWithVersion: false, // Conf according to the version
	}
	if Db, err = gorm.Open(mysql.New(mysqlConfig), logMode(Conf.Mysql.LogMode)); err != nil {
		Logger.Error("Fail to connect to mysql.", zap.Error(err))
		panic(err)
	}
	// set further Conf
	sqlDB, err := Db.DB()
	if err != nil {
		Logger.Error("Fail to get sql db.", zap.Error(err))
		panic(err)
	}
	sqlDB.SetMaxOpenConns(Conf.Mysql.MaxOpenConns)
	sqlDB.SetMaxIdleConns(Conf.Mysql.MaxIdleConns)

	Logger.Info("mysql connected.")
}

// gormConfig dis/enable log according to config
func logMode(mod bool) *gorm.Config {
	var logMode logger.Interface
	if mod {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = logger.Default.LogMode(logger.Silent)
	}
	return &gorm.Config{
		Logger:                                   logMode,
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
