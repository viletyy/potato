/*
 * @Date: 2021-03-11 17:51:14
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-11 18:00:46
 * @FilePath: /hello/util/crypt/rsa.go
 */
package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 私钥生成
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXwIBAAKBgQC1mv515+z3BQYim0nKZPrYm94fPoX0Ew/AJggSdnA2shRsJdMk
AWretmVVej0AH3mBJ1+6tBXXB8+d+7ffNeHWJYrzoh7gPompQWEIvnArcwnhvTA+
8xtdd82GWDKC46pNSqhUGEG/oEqq8Y7C3YoReppJWX7M7xut+YHI+m1hhwIDAQAB
AoGBAIO5WLjM8OR7kGepm2xislBLPmILR74x1UraSyCZJ+uEX6vSA8QqAwpn4jiN
4ZElQ0ya8qTJ2s2NrNo6qrQMsTEaLKYrombAkZbghRYjeAhVCAnIFdchLpV/BG7n
NqxdzV8p1iZeK7qGQm9SWWzS60eO+5RCeUWRB9l3VKai242BAkEA32iXN6bZAIeW
wUIr5GiOV8l6Puw8Or/nRMISHqo4qpbZu2btUTmOuHbLNMo+vVIxCtem2QYGsgMj
9P44fVwJtQJBANAZObuBUkbsqbX/CA+ikDKDtN6wET6S1KaKTU0P5XhxIWeBE8at
EQNqzqaTZ+IMeohjDxTzFlFb3VhHFi2cU8sCQQDM5tSqijDFN5ahMdun5e2HvpaM
V4b2K0Ql4AlWbrECZNDV/JT0xmGL9ghyJnxcj6HDW/7/VXOWmSLFdNTCxUNxAkEA
i8hGBXOKxuV14jBbQ9VYsvXRarwt+TA781p3Lkp9Q3gKjjIgDJZ4FSmLgk0FvMvR
CwgvO01GMoRYnFGzzhNyHQJBALy1gjYLIwHwZYb5ul0s1KJGbY+Tw1qabyQ6lm6A
n5sq4iney5DPUjKcxcP17XvlD9lx7GOS8/ZVWx4krPkydOg=
-----END RSA PRIVATE KEY-----
`)

// 公钥: 根据私钥生成
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC1mv515+z3BQYim0nKZPrYm94f
PoX0Ew/AJggSdnA2shRsJdMkAWretmVVej0AH3mBJ1+6tBXXB8+d+7ffNeHWJYrz
oh7gPompQWEIvnArcwnhvTA+8xtdd82GWDKC46pNSqhUGEG/oEqq8Y7C3YoReppJ
WX7M7xut+YHI+m1hhwIDAQAB
-----END PUBLIC KEY-----
`)

// 加密
// 这里最后的结果需要用 base64.StdEncoding.EncodeToString(data)
func RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
