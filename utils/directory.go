/*
 * @Date: 2021-03-22 10:45:37
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-23 09:29:06
 * @FilePath: /potato/utils/directory.go
 */
package utils

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			fmt.Println("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				fmt.Println("create directory"+v, zap.Any(" error:", err))
			}
		}
	}
	return err
}
