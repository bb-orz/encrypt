package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"encrypt/coding"
)



/*
生成公钥和私钥文件
@param bits RSA bit size
@param pemKeyPath 生成公私钥的路径
*/
func GenerateKey(bits int,pemKeyPath string) error {

	// 生成私钥
	// 使用rsa中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	// 使用pem编码工具生成公私钥pem文件
	err = coding.GenRSAPemKeyFile(privateKey, pemKeyPath)
	if err != nil {
		return err
	}
	return nil

}
