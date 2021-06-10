/*
 * @Date: 2021-06-10 18:51:48
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 18:53:10
 * @FilePath: /potato/internal/service/service.go
 */
package service

import (
	"context"

	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/dao"
)

type Service struct {
	Ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{Ctx: ctx}
	svc.dao = dao.New(global.GO_DB)

	return svc
}
