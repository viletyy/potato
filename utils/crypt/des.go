/*
 * @Date: 2021-03-11 17:25:13
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-11 17:41:39
 * @FilePath: /hello/util/crypt/des.go
 */
package crypt

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
)

/*
 * DES是一种对称加密算法,又称为美国数据加密标准.DES加密时以64位分组对数据进行加密,加密和解密都使用的是同一个长度为64位的密钥,
 * 实际上只用到了其中的56位,密钥中的第8,16…64位用来作奇偶校验.DES有ECB（电子密码本）和CBC（加密块）等加密模式. DES算法的安
 * 全性很高,目前除了穷举搜索破解外, 尚无更好的的办法来破解.其密钥长度越长,破解难度就越大. 填充和去填充函数.
 */

// 示例 text := "I love this beautiful world!"
// 示例 key := []byte("2fa6c1e9")

// Des加密
func DesEncrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

// Des解密
func DesDecrypt(decrypted string, key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)

	return string(out), nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)

	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
