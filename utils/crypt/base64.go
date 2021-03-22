/*
 * @Date: 2021-03-11 18:14:52
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-11 18:19:15
 * @FilePath: /hello/util/crypt/base64.go
 */
package crypt

import "encoding/base64"

/*
 * Base64是一种任意二进制到文本字符串的编码方法,常用于在URL,Cookie,网页中传输少量二进制数据.
 * 首先使用Base64编码需要一个含有64个字符的表,这个表由大小写字母,数字,+和/组成.采用Base64编
 * 码处理数据时,会把每三个字节共24位作为一个处理单元,再分为四组,每组6位,查表后获得相应的字符即
 * 编码后的字符串.编码后的字符串长32位,这样,经Base64编码后,原字符串增长1/3.如果要编码的数据不
 * 是3的倍数,最后会剩下一到两个字节,Base64编码中会采用\x00在处理单元后补全,编码后的字符串最后
 * 会加上一到两个 = 表示补了几个字节.
 */

const (
	base64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"
)

var coder = base64.NewEncoding(base64Table)

func Base64Encode(src []byte) []byte { //编码
	return []byte(coder.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) { //解码
	return coder.DecodeString(string(src))
}
