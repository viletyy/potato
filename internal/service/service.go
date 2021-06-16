/*
 * @Date: 2021-06-10 18:51:48
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-16 23:58:59
 * @FilePath: /potato/internal/service/service.go
 */
package service

import (
	"context"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/dao"
)

type Service struct {
	Ctx context.Context
	Dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{Ctx: ctx}
	svc.Dao = dao.New(otgorm.WithContext(svc.Ctx, global.GO_DB))

	return svc
}
