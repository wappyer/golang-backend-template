package repository

import (
	"context"
	"gitee.com/wappyer/golang-backend-template/internal/domain/model"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/errno"
)

type IAdminRepository interface {
	AutoMigrate(ctx context.Context) error
	Get(ctx context.Context, m *model.Admin) (bool, error)
	MustGet(ctx context.Context, m *model.Admin) errno.Errno
	Add(ctx context.Context, m *model.Admin) error
	AddBatch(ctx context.Context, m []*model.Admin) error
	Update(ctx context.Context, m *model.Admin) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, req AdminListReq) (int64, []*model.Admin, error)
}

type AdminRepository struct {
}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}

func (a *AdminRepository) AutoMigrate(ctx context.Context) error {
	return GetAppDBWithCtx(ctx).AutoMigrate(&model.Admin{})
}

func (a *AdminRepository) Get(ctx context.Context, m *model.Admin) (bool, error) {
	tmp := model.Admin{}
	if *m == tmp { // 空结构体也能查出记录，过滤一下
		return false, nil
	}
	tx := GetAppDBWithCtx(ctx).Where(m).Limit(1).Find(m)
	return tx.RowsAffected > 0, tx.Error
}

func (a *AdminRepository) MustGet(ctx context.Context, m *model.Admin) errno.Errno {
	has, err := a.Get(ctx, m)
	if err != nil {
		return errno.NewErrno(errno.CodeDatabase)
	}
	if !has {
		return errno.NewErrno(errno.CodeLoginNameOrPassword)
	}
	return errno.Errno{}
}

func (a *AdminRepository) Add(ctx context.Context, m *model.Admin) error {
	return GetAppDBWithCtx(ctx).Create(m).Error
}

func (a *AdminRepository) AddBatch(ctx context.Context, m []*model.Admin) error {
	return GetAppDBWithCtx(ctx).Create(m).Error
}

func (a *AdminRepository) Update(ctx context.Context, m *model.Admin) error {
	return GetAppDBWithCtx(ctx).Model(m).Updates(m).Error
}

func (a *AdminRepository) Delete(ctx context.Context, id int) error {
	return GetAppDBWithCtx(ctx).Delete(&model.Admin{}, id).Error
}

type AdminListReq struct {
	Search string
}

func (a *AdminRepository) List(ctx context.Context, req AdminListReq) (int64, []*model.Admin, error) {
	var list []*model.Admin
	var count int64
	err := GetAppDBWithCtx(ctx).Where("`name` like ?", "%"+req.Search+"%").Find(&list).Count(&count).Error
	return count, list, err
}
