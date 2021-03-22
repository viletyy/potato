/*
 * @Date: 2021-03-22 17:50:15
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 17:53:11
 * @FilePath: /potato/global/model.go
 */
package global

import "time"

type Model struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}
