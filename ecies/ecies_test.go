package ecies

import (
	"testing"
)

const PemKeyPath = "./"

func TestECIES(t *testing.T) {
	// 获取随机字符串
	randomKey := GetRandomString(40)

	// 生成私钥和公钥
	err := GenerateKey(randomKey,PemKeyPath)
	if err != nil {
		t.Fatal(err.Error())
	}

	// 加密前源信息
	srcInfo := "ECIES 椭圆曲线实现数据加解密测试"
	t.Log("原文：", srcInfo)

	// 加密信息
	cryptData, err := EnCryptByEcies(srcInfo, PemKeyPath+"ECCPublicKey.pem")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("ECIES加密后为：", cryptData)

	// 解密信息
	srcData, err := DeCryptByEcies(cryptData, PemKeyPath+"ECCPrivateKey.pem")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("ECIES解密后为：", srcData)

	if srcInfo == string(srcData) {
		t.Log("ECIES加解密验证成功")
	} else {
		t.Error("ECIES加解密验证失败")
	}
}
