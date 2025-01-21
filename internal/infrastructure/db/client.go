package db

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/db/model"
	"gorm.io/gorm"
	"sync"
)

const ContextKeyGormDbVar = "gormDBVar"

type Client struct {
	dbConf   Config
	rw       sync.RWMutex
	globalDB *gorm.DB
}

func NewDBClient(linkConf Config) (*Client, error) {
	d := &Client{
		dbConf: linkConf,
	}
	db, err := d.createDb()
	if err != nil {
		return d, err
	}
	d.globalDB = db

	return d, nil
}

func (d *Client) GetDB() *gorm.DB {
	d.rw.RLock()
	defer d.rw.RUnlock()
	return d.globalDB
}

func (d *Client) createDb() (*gorm.DB, error) {
	if d.dbConf.DbType == DSNMysql {
		return gormMysql(d.dbConf)
	}
	return nil, errors.New("db type not support")
}

func (d *Client) updateDb(linkConf Config) error {
	if linkConf.DbType == DSNMysql {
		return updateMysqlConn(linkConf, d.globalDB)
	}
	return errors.New("db type not support")
}

func (d *Client) CreateView() error {
	for k, v := range model.SqlMap {
		err := d.globalDB.Exec(v).Error
		if err != nil {
			return errors.New(fmt.Sprintf("视图创建%v失败：%v", k, err))
		}
	}
	return nil
}

func WithValueAppDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, ContextKeyGormDbVar, db)
}

func (d *Client) GetAppDBWithCtx(ctx context.Context) *gorm.DB {
	return d.GetAppDB(ctx).WithContext(ctx)
}

func (d *Client) GetAppDB(ctx context.Context) *gorm.DB {
	value := ctx.Value(ContextKeyGormDbVar)
	if value != nil {
		gDB, ok := value.(*gorm.DB)
		if ok {
			return gDB
		}
	}
	return d.GetDB()
}

func (d *Client) DbTransaction(ctx context.Context, fn func(context.Context) error) error {
	return d.GetAppDB(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(WithValueAppDB(ctx, tx))
	})
}
