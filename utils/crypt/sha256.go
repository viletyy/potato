/*
 * @Date: 2021-03-11 18:12:11
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-11 18:14:43
 * @FilePath: /hello/util/crypt/sha256.go
 */
package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func Sha256Check(content, secret, encrypted string) bool {
	return strings.EqualFold(Sha256Encode(content, secret), encrypted)
}

func Sha256Encode(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
