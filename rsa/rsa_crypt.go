package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/gofuncchan/encrypt/coding"
)


/*
RSA算法公钥加密
@param src 源数据字节
@param pubKeyFile 公钥文件
*/
func EnCryptByRSA(src []byte, pubKeyFile string) ([]byte, error) {
	// 从publicKey.pem解析公钥内容
	publicKey, err := coding.ParseRSAPublicPemKey(pubKeyFile)
	if err != nil {
		return nil, err
	}

	// 使用得到的公钥通过rsa进行数据加密
	encryptBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, src)
	if err != nil {
		return nil, err
	}
	return encryptBytes, nil
}


/*
RSA算法私钥解密
@param cryptBytes 密文字节数据
@param privateKeyFile 私钥文件
*/
func DeCryptByRSA(cryptBytes []byte, privateKeyFile string) ([]byte, error) {
	// 从privateKey.pem文件解析私钥内容
	privateKey, err := coding.ParseRSAPrivatePemKey(privateKeyFile)

	// 使用得到的私钥通过rsa进行数据解密
	decryptBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cryptBytes)
	if err != nil {
		return nil, err
	}
	return decryptBytes, nil
}
