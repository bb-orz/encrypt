package ecies

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"encrypt/coding"
)


/*
ECIES 公钥数据加密
@param srcData 源数据
@param publicFile 公钥文件
*/
func EnCryptByEcies(srcData, pubKeyFile string) (cryptData string, err error) {
	// 获取公钥数据
	publicKey, err := coding.ParseECCPublicPemKey(pubKeyFile)
	if err != nil {
		return "", err
	}

	// 转换为ecies所需要的公钥格式
	publicKeyForEcies := ecies.ImportECDSAPublic(publicKey)


	// 公钥加密数据
	encryptBytes, err := ecies.Encrypt(rand.Reader, publicKeyForEcies, []byte(srcData), nil, nil)
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
func DeCryptByEcies(cryptData, privateKeyFile string) (srcData string, err error) {
	// 获取私钥信息
	privateKey, err := coding.ParseECCPrivatePemKey(privateKeyFile)
	if err != nil {
		return "", err
	}

	// 转换为ecies所需要的私钥格式
	privateKeyForEcies := ecies.ImportECDSA(privateKey)

	// 私钥解密数据
	cryptBytes, err := hex.DecodeString(cryptData)
	srcByte, err := privateKeyForEcies.Decrypt(cryptBytes, nil, nil)
	if err != nil {
		fmt.Println("解密错误：", err)
		return "", err
	}
	srcData = string(srcByte)

	return
}
