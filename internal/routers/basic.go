/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-11 00:22:23
 * @FilePath: /potato/internal/routers/basic.go
 */
package routers

import "github.com/viletyy/potato/internal/controller/api/v1/basic"

func V1InitBasicRouter() {

	vendors := V1RouterGroup.Group("/vendors")
	vendor := basic.NewVendor()
	{
		vendors.GET("", vendor.List)
		vendors.POST("", vendor.Create)
		vendors.GET("/:id", vendor.Get)
		vendors.PATCH("/:id", vendor.Update)
		vendors.DELETE("/:id", vendor.Delete)
	}
}
