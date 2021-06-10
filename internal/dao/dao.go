/*
 * @Date: 2021-06-10 18:46:10
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 18:48:41
 * @FilePath: /potato/internal/dao/auth.go
 */
package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	Engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{Engine: engine}
}
