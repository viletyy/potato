/*
 * @Date: 2021-03-11 18:06:18
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-11 18:18:08
 * @FilePath: /hello/util/crypt/md5.go
 */
package crypt

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

/*
 * MD5的全称是Message-DigestAlgorithm 5,它可以把一个任意长度的字节数组转换成一个定长的整数,并且这种转换是不可
 * 逆的.对于任意长度的数据,转换后的MD5值长度是固定的,而且MD5的转换操作很容易,只要原数据有一点点改动,转换后结果就
 * 会有很大的差异.正是由于MD5算法的这些特性,它经常用于对于一段信息产生信息摘要,以防止其被篡改.其还广泛就于操作系
 * 统的登录过程中的安全验证,比如Unix操作系统的密码就是经过MD5加密后存储到文件系统中,当用户登录时输入密码后, 对用
 * 户输入的数据经过MD5加密后与原来存储的密文信息比对,如果相同说明密码正确,否则输入的密码就是错误的. MD5以512位为
 * 一个计算单位对数据进行分组,每一分组又被划分为16个32位的小组,经过一系列处理后,输出4个32位的小组,最后组成一个128
 * 位的哈希值.对处理的数据进行512求余得到N和一个余数,如果余数不为448,填充1和若干个0直到448位为止,最后再加上一个64
 * 位用来保存数据的长度,这样经过预处理后,数据变成（N+1）x 512位. 加密.Encode 函数用来加密数据,Check函数传入一个
 * 未加密的字符串和与加密后的数据,进行对比,如果正确就返回true.
 */

func Md5Check(content, encrypted string) bool {
	return strings.EqualFold(Md5Encode(content), encrypted)
}
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
