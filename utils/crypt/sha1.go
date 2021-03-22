/*
 * @Date: 2021-03-11 18:07:34
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-11 18:11:49
 * @FilePath: /hello/util/crypt/sha1.go
 */
package crypt

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func Sha1Check(content, encrpyted string) bool {
	return strings.EqualFold(Sha1Encode(content), encrpyted)
}

func Sha1Encode(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
