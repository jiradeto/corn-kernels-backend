package connectors

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// postgresDBInstance is an instance for of postgres db in gorm format
	postgresDBInstance *gorm.DB
	postgresDBLogMode  bool
)

// ConnectSqliteDB is a connector function for connecting postgres
func ConnectSqliteDB() *gorm.DB {
	return connect(postgresDBLogMode)
}

func connect(dbLogMode bool) *gorm.DB {
	var err error
	mut.Lock()
	defer mut.Unlock()
	var logLevel logger.LogLevel
	if dbLogMode {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}
	if postgresDBInstance == nil {
		postgresDBInstance, err = gorm.Open(sqlite.Open("local.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})
		if err != nil {
			panic(err.Error())
		}
	}
	return postgresDBInstance
}
