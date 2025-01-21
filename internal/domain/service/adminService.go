package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"gitee.com/wappyer/golang-backend-template/internal/domain/entity"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/db/model"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/db/repository"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/errno"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/jwt"
)

type AdminService struct {
	AdminRepo repository.IAdminRepository
}

func NewAdminService() *AdminService {
	return &AdminService{
		AdminRepo: repository.NewAdminRepository(),
	}
}

func (s *AdminService) Admin(ctx context.Context, phone, password string) (string, entity.Admin, errno.Errno) {
	pwdMd5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	admin := &model.Admin{Phone: phone, Password: pwdMd5}
	adminEntity := entity.Admin{Admin: admin}

	// 验证帐号密码
	eno := s.AdminRepo.MustGet(ctx, adminEntity.Admin)
	if eno.NotNil() {
		return "", adminEntity, eno
	}

	// 生成token
	token, err := jwt.GetClientIns().GenerateToken(admin.Id)
	if err != nil {
		return "", adminEntity, errno.NewErrno(errno.CodeCreateToken)
	}

	return token, adminEntity, errno.Errno{}
}
