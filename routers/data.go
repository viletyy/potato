package routers

import (
	"github.com/viletyy/potato/controller/api/v1/data"
)

func V1InitDataRouter() {
	metaDatabases := V1RouterGroup.Group("/meta_databases")
	{
		metaDatabases.GET("", data.GetMetaDatabases)
		metaDatabases.GET(":id/meta_tables", data.GetMetaTables)
	}
}
