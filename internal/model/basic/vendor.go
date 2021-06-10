/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 17:54:47
 * @FilePath: /potato/internal/model/basic/vendor.go
 */
package basic

import (
	"github.com/viletyy/potato/internal/model"
)

type Vendor struct {
	model.Model

	Name string `json:"name"`
	Uuid int    `json:"uuid"`
}
