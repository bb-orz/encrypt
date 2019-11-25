package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)



/*
生成公钥和私钥文件
@param bits RSA bit size
@param pemKeyPath 生成公私钥的路径
*/
func GenerateKey(bits int,pemKeyPath string) error {
	var privateKeyFile = pemKeyPath + "/privateKey.pem"
	var publicKeyFile = pemKeyPath + "/publicKey.pem"

	// 生成私钥
	// 使用rsa中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	// 通过x509标准将得到的ras私钥序列化为ASN.1的DER编码字符串
	PKCS1PrivateBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	// 将私钥字符串设置到pem格式块中
	privateBlock := pem.Block{
		Type:  "RSA Private Key",
		Bytes: PKCS1PrivateBytes,
	}

	// 通过pem将设置好的数据进行编码，并写入磁盘文件
	privateFile, err := os.Create(privateKeyFile)
	if err != nil {
		return err
	}
	defer privateFile.Close()
	err = pem.Encode(privateFile, &privateBlock)
	if err != nil {
		return err
	}

	// 生成公钥
	// 从得到的私钥对象中将公钥信息取出
	publicKey := privateKey.PublicKey

	// 通过x509标准将得到的ras公钥序列化为ASN.1的DER编码字符串
	PKCS1PublicBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}

	// 将公钥字符串设置到pem格式块中
	publicBlock := pem.Block{
		Type:  "RSA Public Key",
		Bytes: PKCS1PublicBytes,
	}

	// 通过pem将设置好的数据进行编码，并写入磁盘文件
	publicFile, err := os.Create(publicKeyFile)
	if err != nil {
		return err
	}
	defer publicFile.Close()
	err = pem.Encode(publicFile, &publicBlock)
	if err != nil {
		return err
	}

	return nil
}
