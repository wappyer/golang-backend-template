package entity

import "gitee.com/wappyer/golang-backend-template/internal/domain/model"

type Log struct {
	*model.Log
	Detail *model.LogDetail
}
