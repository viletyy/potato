/*
 * @Date: 2021-06-11 17:27:25
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-11 17:54:34
 * @FilePath: /potato/internal/controller/api/v1/upload.go
 */
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/service"
	"github.com/viletyy/potato/pkg/app"
	"github.com/viletyy/potato/pkg/errcode"
	"github.com/viletyy/potato/pkg/upload"
	"github.com/viletyy/yolk/convert"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// @Summary 上传文件
// @Description
// @Accept mpfd
// @Produce json
// @Param token header string true "auth by /auth"
// @Param file formData file true "文件"
// @Param type formData int true "文件类型"
// @Success 200 {object} errcode.Error "请求成功"
// @Router /v1/upload [post]
func (u Upload) Create(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.GO_LOG.Sugar().Errorf("c.Request.FormFile err: %v", err)
		errorCode := errcode.InvalidParams
		errorCode.WithData(err)
		response.ToErrorResponse(errorCode)
		return
	}

	fileType, err := convert.StrTo(c.PostForm("type")).Int()
	if err != nil {
		global.GO_LOG.Sugar().Errorf("convert.StrTo err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams)
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.UploadFile err :v", err)
		errorCode := errcode.ErrorUploadFileFail
		errorCode.WithData(err)
		response.ToErrorResponse(errorCode)
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
