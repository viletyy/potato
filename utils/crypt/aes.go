/*
 * @Date: 2021-03-11 17:03:16
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-11 17:42:16
 * @FilePath: /hello/util/crypt/aes.go
 */
package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/beego/beego/v2/core/logs"
)

/*
 * AES,即高级加密标准（Advanced Encryption Standard）,是一个对称分组密码算法,旨在取代DES成为广泛使用的标准.
 * AES中常见的有三种解决方案,分别为AES-128,AES-192和AES-256. AES加密过程涉及到4种操作：字节替代（SubBytes）
 * ,行移位（ShiftRows）,列混淆（MixColumns）和轮密钥加（AddRoundKey）.解密过程分别为对应的逆操作.由于每一步
 * 操作都是可逆的,按照相反的顺序进行解密即可恢复明文.加解密中每轮的密钥分别由初始密钥扩展得到.算法中16字节的明文,
 * 密文和轮密钥都以一个4x4的矩阵表示. AES 有五种加密模式：电码本模式（Electronic Codebook Book (ECB)）,密
 * 码分组链接模式（Cipher Block Chaining (CBC)）,计算器模式（Counter (CTR)）,密码反馈模式（Cipher Feed
 * Back (CFB)）和输出反馈模式（Output FeedBack (OFB)）
 */

// 示例 orig := "hello world"
// 示例 key := "123456781234567812345678"

//AES加密
func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组密钥
	block, err := aes.NewCipher(k)
	if err != nil {
		logs.Error(fmt.Sprintf("key 长度必须 16/24/32长度：%s", err.Error()))
	}
	// 获取密钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	// 使用RawURLEncoding 不要使用StdEncoding
	// 不要使用StdEncoding 放在url参数中会导致错误
	return base64.RawURLEncoding.EncodeToString(cryted)
}

//AES解密
func AesDecrypt(cryted string, key string) string {
	// 使用RawURLEncoding 不要使用StdEncoding
	// 不要使用StdEncoding  放在url参数中回导致错误
	crytedByte, _ := base64.RawURLEncoding.DecodeString(cryted)
	k := []byte(key)

	// 分组密钥
	block, err := aes.NewCipher(k)
	if err != nil {
		logs.Error(fmt.Sprintf("key 长度必须 16/24/32长度：%s", err.Error()))
	}
	// 获取密钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)

	return string(orig)
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
