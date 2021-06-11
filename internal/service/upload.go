/*
 * @Date: 2021-06-11 17:07:48
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-11 17:24:45
 * @FilePath: /potato/internal/service/upload.go
 */
package service

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if !upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}

	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.GO_CONFIG.App.UploadServerPath + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
