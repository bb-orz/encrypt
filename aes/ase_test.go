package aes

import (
	"encoding/hex"
	"testing"
)

const (
	MYKEY = "abcdefgh12345678"
	IV    = "aaaabbbb12345678"
)

func TestAES(t *testing.T) {
	t.Log("测试AES对称加解密：")
	srcString := "AES对称加解密测试"
	// 加密
	cryptBytes,err := EncryptAES([]byte(srcString), []byte(MYKEY),[]byte(IV))
	if err != nil{
		t.Error(err.Error())
	}
	t.Log("加密效果：")
	t.Log("src :", hex.EncodeToString([]byte(srcString)), "[]byte长度:", len([]byte(srcString)))
	t.Log("dst :", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))

	// 解密
	deCryptBytes,err := DecryptAES(cryptBytes, []byte(MYKEY),[]byte(IV))
	if err != nil{
		t.Error(err.Error())
	}
	t.Log("解密效果：")
	t.Log("src bytes:", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))
	t.Log("dst bytes:", hex.EncodeToString(deCryptBytes), "[]byte长度:", len(deCryptBytes))

	t.Log("原字串为：" + srcString)
	t.Log("解密字符为：" + string(deCryptBytes))
	if srcString == string(deCryptBytes) {
		t.Log("AES加解密验证成功")
	} else {
		t.Error("AES加解密验证失败")
	}
}
