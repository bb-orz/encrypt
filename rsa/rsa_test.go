package rsa

import (
	"encoding/hex"
	"testing"
)

const PemKeyPath = "./"

func TestRSA(t *testing.T) {
	// 生成钥匙对
	err := GenerateKey(4096, PemKeyPath)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("生成RSA公私钥成功，生成路径：", PemKeyPath)

	srcInfo := "RSA非对称加解密测试"

	// 公钥加密
	encryptBytes, err := EnCryptByRSA([]byte(srcInfo), PemKeyPath+"publicKey.pem")
	if err != nil {
		t.Fatal(err)
	}

	// 私钥解密
	decryptBytes, err := DeCryptByRSA(encryptBytes, PemKeyPath+"privateKey.pem")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("测试非对称加密结果：")
	t.Log("源数据：", srcInfo)
	encryptHex := hex.EncodeToString(encryptBytes)
	t.Log("公钥加密数据：", encryptHex)
	t.Log("私钥解密数据：", string(decryptBytes))

	if srcInfo == string(decryptBytes) {
		t.Log("RSA加解密验证成功")
	} else {
		t.Error("RSA加解密验证失败")
	}

}
