package RSA

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"main.go/tuuz/Base64"
)

func EncB64(publicKey string, origData []byte) (string, error) {
	ret, err := Encrypt(publicKey, origData)
	if err != nil {
		return "", err
	} else {
		return Base64.Encode(ret), nil
	}
}

func Encrypt(publicKey string, origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(publicKey))
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

func Decrypt(privateKey string, ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode([]byte(privateKey))
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

func DecB64(privateKey string, base64string string) ([]byte, error) {
	b64, err := Base64.Decode(base64string)
	if err != nil {
		return b64, err
	}
	return Decrypt(privateKey, b64)
}
