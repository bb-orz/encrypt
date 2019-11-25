package ecies

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)


/*
ECIES 公钥数据加密
@param srcData 源数据
@param publicFile 公钥文件
*/
func EnCryptByEcies(srcData, publicFile string) (cryptData string, err error) {
	// 获取公钥数据
	publicKey, err := GetPublicKeyByPemFile(publicFile)
	if err != nil {
		return "", err
	}

	// 公钥加密数据
	encryptBytes, err := ecies.Encrypt(rand.Reader, publicKey, []byte(srcData), nil, nil)
	if err != nil {
		return "", err
	}

	cryptData = hex.EncodeToString(encryptBytes)

	return
}


/*
ECIES 私钥数据解密
@param cryptData 密文字节数据
@param privateFile 私钥文件
*/
func DeCryptByEcies(cryptData, privateFile string) (srcData string, err error) {
	// 获取私钥信息
	privateKey, err := GetPrivateKeyByPemFile(privateFile)
	if err != nil {
		return "", err
	}

	// 私钥解密数据
	cryptBytes, err := hex.DecodeString(cryptData)
	srcByte, err := privateKey.Decrypt(cryptBytes, nil, nil)
	if err != nil {
		fmt.Println("解密错误：", err)
		return "", err
	}
	srcData = string(srcByte)

	return
}
