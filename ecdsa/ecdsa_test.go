package ecdsa

import (
	"testing"
)
const PemKeyPath = "./"

func TestECDSA(t *testing.T) {
	// 生成随机钥字符串长度40字节，用于生产公私钥证书
	randKey := GetRandomString(40)
	// 生成随机签名字符串40字节，用于加密数据
	randSign := GetRandomString(40)

	// 使用随机钥字符串生成公私钥文件
	err := GenerateKey(randKey,PemKeyPath)
	if err != nil {
		t.Fatal(err.Error())
	}

	// 签名附加信息
	srcInfo := "ECDSA 椭圆曲线实现数字签名"
	t.Log("原文：", srcInfo)

	// ECC签名加密
	signByEcc, err := CryptSignByEcc(srcInfo, PemKeyPath+"privateKey.pem", randSign)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("ECDSA私钥加密签名为：", signByEcc)

	// ECC签名算法校验
	verifyCryptEcc, err := VerifyCryptEcc(srcInfo, signByEcc,PemKeyPath+"publicKey.pem")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("ECDSA公钥解密后验签校验结果：", verifyCryptEcc)
}
