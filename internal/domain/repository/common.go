package repository

import (
	"context"
	"gitee.com/wappyer/golang-backend-template/config"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/db"
	"gorm.io/gorm"
)

const ContextKeyGormDbVar = "gormDBVar"

type IRepo interface {
	AutoMigrate(ctx context.Context) error
}

type IDBFactory interface {
	GetDB() *gorm.DB
}

type RepoDBFactory struct {
	Repos     []IRepo
	DBFactory IDBFactory
}

var repoDBFactory = &RepoDBFactory{}

func Initialize(cfg config.DB) error {
	var err error
	repoDBFactory.DBFactory, err = db.NewDBFactory(cfg)
	if err != nil {
		return err
	}

	repoDBFactory.Repos = []IRepo{
		NewLogRepository(),
		NewAdminRepository(),
	}

	ctx := WithValueAppDB(context.Background(), repoDBFactory.DBFactory.GetDB())
	for _, repo := range repoDBFactory.Repos {
		err = repo.AutoMigrate(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func WithValueAppDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, ContextKeyGormDbVar, db)
}

func GetAppDB(ctx context.Context) *gorm.DB {
	value := ctx.Value(ContextKeyGormDbVar)
	if value != nil {
		gDB, ok := value.(*gorm.DB)
		if ok {
			return gDB
		}
	}
	return repoDBFactory.DBFactory.GetDB()
}

func DbTransaction(ctx context.Context, fn func(context.Context) error) error {
	return GetAppDB(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(WithValueAppDB(ctx, tx))
	})
}

func GetAppDBWithCtx(ctx context.Context) *gorm.DB {
	return GetAppDB(ctx).WithContext(ctx)
}
