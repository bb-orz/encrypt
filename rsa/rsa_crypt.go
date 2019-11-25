package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)


/*
RSA算法公钥加密
@param src 源数据字节
@param pubKeyFile 公钥文件
*/
func EnCryptByRSA(src []byte, pubKeyFile string) ([]byte, error) {
	var err error
	// 将公钥文件中的公钥读出，得到使用pem编码的字符串
	file, err := os.Open(pubKeyFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, fileInfo.Size())
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	// 将得到的字符串解码
	block, _ := pem.Decode(buffer)

	// 使用x509将编码之后的公钥解析出来
	pubInner, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey := pubInner.(*rsa.PublicKey)

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
	// 将私钥文件中的私钥读出，得到使用pem编码的字符串
	file, err := os.Open(privateKeyFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := fileInfo.Size()
	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	// 将得到的字符串解码
	block, _ := pem.Decode(buffer)

	// 使用x509将编码之后的私钥解析出来
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 使用得到的私钥通过rsa进行数据解密
	decryptBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cryptBytes)
	if err != nil {
		return nil, err
	}
	return decryptBytes, nil
}
