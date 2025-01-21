package entity

import (
	model "gitee.com/wappyer/golang-backend-template/internal/infrastructure/db/model"
)

type Log struct {
	*model.Log
	Detail *model.LogDetail
}
