/*
 * @Date: 2021-06-10 16:47:58
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 22:28:15
 * @FilePath: /potato/pkg/app/app.go
 */
package app

import (
	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	err := errcode.Success
	err.WithData(data)

	r.ToErrorResponse(err)
}

func (r *Response) ToResponseErrors(data interface{}) {
	err := errcode.InvalidParams
	err.WithData(data)
	r.ToErrorResponse(err)
}

func (r *Response) ToResponseList(list interface{}, total int) {
	err := errcode.Success
	err.WithData(map[string]interface{}{
		"list": list,
		"pager": Pager{
			Page:     GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			Total:    total,
		},
	})
	r.ToErrorResponse(err)
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code, "msg": err.Msg}
	data := err.Data
	response["data"] = data

	r.Ctx.JSON(err.StatusCode(), response)
}
