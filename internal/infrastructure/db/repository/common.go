package repository

import (
	"context"
	"gitee.com/wappyer/golang-backend-template/config"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/db"
)

type IRepo interface {
	AutoMigrate(ctx context.Context) error
}

type RepoFactory struct {
	Repos []IRepo
}

var repoFactory = &RepoFactory{}
var DB *db.Client

func Initialize(cfg config.DB) error {
	var err error
	dbConfig := db.Config{
		DbType:      cfg.DbType,      // 数据库类型 mysql
		Path:        cfg.Path,        // 服务器地址
		Dbname:      cfg.Dbname,      // 数据库名
		Config:      cfg.Config,      // 高级配置
		Username:    cfg.Username,    // 登录名
		Password:    cfg.Password,    // 密码
		MaxIdleConn: cfg.MaxIdleConn, // 空闲中的最大连接数
		MaxOpenConn: cfg.MaxOpenConn, // 打开到数据库的最大连接数
		RedisEnable: cfg.RedisEnable, // 是否开启redis
	}
	DB, err = db.NewDBClient(dbConfig)
	if err != nil {
		return err
	}

	ctx := db.WithValueAppDB(context.Background(), DB.GetDB())
	for _, repo := range repoFactory.Repos {
		err = repo.AutoMigrate(ctx)
		if err != nil {
			return err
		}
	}

	// 创建视图
	_ = DB.CreateView()

	return nil
}
