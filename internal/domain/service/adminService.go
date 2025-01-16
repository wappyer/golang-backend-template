package service

import (
	"context"
	"gitee.com/wappyer/golang-backend-template/internal/domain/model"
	"gitee.com/wappyer/golang-backend-template/internal/domain/repository"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/errno"
)

type AdminService struct {
	AdminRepo repository.IAdminRepository
}

func NewAdminService() *AdminService {
	return &AdminService{
		AdminRepo: repository.NewAdminRepository(),
	}
}

func (s *AdminService) Admin(ctx context.Context, phone, password string) (string, errno.Errno) {
	admin := &model.Admin{Phone: phone}

	// 验证帐号密码
	eno := s.AdminRepo.MustGet(ctx, admin)
	if eno.NotNil() {
		return "", eno
	}
	if admin.Password != password {
		return "", errno.NewErrno(errno.CodeAdminLoginPassword)
	}

	return admin.Phone, errno.Errno{}
}
