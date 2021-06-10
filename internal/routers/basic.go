/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 17:51:43
 * @FilePath: /potato/internal/routers/basic.go
 */
package routers

func V1InitBasicRouter() {

	vendors := V1RouterGroup.Group("/vendors")
	{
		vendors.GET("")
		vendors.POST("")
		vendors.PATCH("/:id")
		vendors.DELETE("/:id")
	}
}
