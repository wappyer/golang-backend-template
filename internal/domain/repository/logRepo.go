package repository

import (
	"context"
	"gitee.com/wappyer/golang-backend-template/internal/domain/entity"
	"gitee.com/wappyer/golang-backend-template/internal/domain/model"
)

type ILogRepository interface {
	Insert(ctx context.Context, log *entity.Log) error
	Update(ctx context.Context, log *entity.Log) error
}

type LogRepository struct{}

func NewLogRepository() *LogRepository {
	return &LogRepository{}
}

func (a *LogRepository) AutoMigrate(ctx context.Context) error {
	return GetAppDBWithCtx(ctx).AutoMigrate(&model.Log{}, &model.LogDetail{})
}

func (a *LogRepository) Insert(ctx context.Context, log *entity.Log) error {
	return DbTransaction(ctx, func(ctx context.Context) error {
		err := GetAppDBWithCtx(ctx).Create(log.Log).Error
		if err != nil {
			return err
		}
		return GetAppDBWithCtx(ctx).Create(log.Detail).Error
	})
}

func (a *LogRepository) Update(ctx context.Context, log *entity.Log) error {
	return DbTransaction(ctx, func(ctx context.Context) error {
		err := GetAppDBWithCtx(ctx).Model(&model.Log{}).Where("request_id = ?", log.Log.RequestId).Updates(log.Log).Error
		if err != nil {
			return err
		}
		return GetAppDBWithCtx(ctx).Model(&model.LogDetail{}).Where("request_id = ?", log.Log.RequestId).Updates(log.Detail).Error
	})
}
