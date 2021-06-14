/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-14 23:32:56
 * @FilePath: /potato/internal/routers/basic.go
 */
package routers

import "github.com/viletyy/potato/internal/controller/api/v1/basic"

func V1InitBasicRouter() {
	v1RouterGroup := Engine.Group("../api/v1")

	vendors := v1RouterGroup.Group("/vendors")
	vendor := basic.NewVendor()
	{
		vendors.GET("", vendor.List)
		vendors.POST("", vendor.Create)
		vendors.GET("/:id", vendor.Get)
		vendors.PATCH("/:id", vendor.Update)
		vendors.DELETE("/:id", vendor.Delete)
	}
}
