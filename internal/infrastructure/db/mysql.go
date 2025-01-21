package db

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DSNMysql = "mysql"

func gormMysql(dbc Config) (db *gorm.DB, err error) {
	mysqlConfig := mysql.Config{
		DSN:                       mySqlDsn(dbc),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	if db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return nil, errors.New("mysql connect error")
	}
	return db, updateMysqlConn(dbc, db)
}

func updateMysqlConn(dbc Config, db *gorm.DB) error {
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(dbc.MaxIdleConn)
		sqlDB.SetMaxOpenConns(dbc.MaxOpenConn)
		return nil
	} else {
		return errors.New("mysql connect error")
	}
}

func mySqlDsn(m Config) string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}
