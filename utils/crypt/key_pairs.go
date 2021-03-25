/*
 * @Date: 2021-03-11 18:02:34
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-25 14:47:39
 * @FilePath: /potato/utils/crypt/key_pairs.go
 */
package crypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/viletyy/potato/global"
	"go.uber.org/zap"
)

// ecdsaCmd represents the doc command
func KeyPairs(keyName string) {
	//elliptic.P256(),elliptic.P384(),elliptic.P521()

	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		global.GO_LOG.Error("Genera KeyPairs Error:", zap.Any("err", err))
	}
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	privateBs := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	privateFile, err := os.Create(keyName + ".private.pem")
	if err != nil {
		global.GO_LOG.Error("Genera KeyPairs Error:", zap.Any("err", err))
	}
	_, err = privateFile.Write(privateBs)
	if err != nil {
		global.GO_LOG.Error("Genera KeyPairs Error:", zap.Any("err", err))
	}
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(privateKey.Public())
	publicBs := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	publicKeyFile, err := os.Create(keyName + ".public.pem")
	if err != nil {
		global.GO_LOG.Error("Genera KeyPairs Error:", zap.Any("err", err))
	}
	_, err = publicKeyFile.Write(publicBs)
	if err != nil {
		global.GO_LOG.Error("Genera KeyPairs Error:", zap.Any("err", err))
	}
}
