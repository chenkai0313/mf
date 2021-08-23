package app

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"server/config"
)

var MysqlDb *gorm.DB

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	if config.GetDebug() {
		fmt.Printf(format+"\n", args...)
	} else {
		var lineNum, sqlTime, rows, sql interface{}
		for k, v := range args {
			switch k {
			case 0:
				lineNum = v
			case 1:
				sqlTime = v
			case 2:
				rows = v
			case 3:
				sql = v
			}
		}
		ZapLog.dbLog(lineNum, sqlTime, rows, sql)
	}
}

func InitMysqlDb() {

	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Silent,          // Log level
			IgnoreRecordNotFoundError: false,                  // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(config.GetDb()), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true, //禁止自动创建外建
	})
	if err != nil {
		panic(fmt.Sprintf("mysql connect error= %v", err.Error()))
	}
	if config.GetDebug() {
		db = db.Debug()
	}
	MysqlDb = db
}
