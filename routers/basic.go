/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-24 11:23:01
 * @FilePath: /potato/routers/basic.go
 */
package routers

import (
	"github.com/viletyy/potato/controller/api/v1/basic"
)

func V1InitBasicRouter() {
	// metaDatabases := V1RouterGroup.Group("/meta_databases")
	// {
	// 	metaDatabases.GET("", basic.GetMetaDatabases)
	// 	metaDatabases.POST("", basic.AddMetaDatabase)
	// 	metaDatabases.PATCH("/:id", basic.EditMetaDatabase)
	// 	metaDatabases.DELETE("/:id", basic.DeleteMetaDatabase)
	// 	metaDatabases.GET("/:id/meta_tables", basic.GetMetaTables)
	// }
	vendors := V1RouterGroup.Group("/vendors")
	{
		vendors.GET("", basic.GetVendors)
		vendors.POST("", basic.CreateVendor)
		vendors.PATCH("/:id", basic.UpdateVendor)
		vendors.DELETE("/:id", basic.DeleteVendor)
	}
	// businesses := V1RouterGroup.Group("/businesses")
	// {
	// 	businesses.GET("", basic.GetBusinesses)
	// 	businesses.POST("", basic.AddBusiness)
	// 	businesses.PATCH("/:id", basic.EditBusiness)
	// 	businesses.DELETE("/:id", basic.DeleteBusiness)
	// }
}
