package db

import (
	"errors"
	"fmt"
	"gitee.com/wappyer/golang-backend-template/config"
	"gitee.com/wappyer/golang-backend-template/internal/domain/model/view"
	"gorm.io/gorm"
	"sync"
)

type Factory struct {
	dbConf   config.DB
	rw       sync.RWMutex
	globalDB *gorm.DB
}

func NewDBFactory(linkConf config.DB) (*Factory, error) {
	d := &Factory{}
	db, err := d.createDb(linkConf)
	if err != nil {
		return d, err
	}
	d.globalDB = db
	d.dbConf = linkConf

	// 创建视图
	_ = d.CreateView()
	return d, nil
}

func (d *Factory) GetDB() *gorm.DB {
	d.rw.RLock()
	defer d.rw.RUnlock()
	return d.globalDB
}

func (d *Factory) createDb(linkConf config.DB) (*gorm.DB, error) {
	if linkConf.DbType == "mysql" {
		return gormMysql(linkConf)
	}
	return nil, errors.New("db type not support")
}

func (d *Factory) updateDb(linkConf config.DB) error {
	if linkConf.DbType == "mysql" {
		return updateMysqlConn(linkConf, d.globalDB)
	}
	return errors.New("db type not support")
}

func (d *Factory) CreateView() error {
	for k, v := range view.SqlMap {
		err := d.globalDB.Exec(v).Error
		if err != nil {
			return errors.New(fmt.Sprintf("视图创建%v失败：%v", k, err))
		}
	}
	return nil
}
