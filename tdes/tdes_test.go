package tdes

import (
	"encoding/hex"
	"testing"
)
const (
	MYKEY = "abcdefgh12345678ABCDEFGH" //三组八字节密匙，即24字节
	IV    = "aaaabbbb"                 //CBC模式的初始化向量，8字节
)

func TestTDES(t *testing.T) {
	t.Log("测试TDES对称加解密：")
	srcString := "TDES对称加解密测试"
	// 加密
	cryptBytes,err := EncryptTDES([]byte(srcString), []byte(MYKEY),[]byte(IV))
	if err != nil{
		t.Error(err.Error())
	}
	t.Log("加密效果：")
	t.Log("src :", hex.EncodeToString([]byte(srcString)), "[]byte长度:", len([]byte(srcString)))
	t.Log("dst :", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))

	// 解密
	deCryptBytes,err := DecryptTDES(cryptBytes, []byte(MYKEY),[]byte(IV))
	if err != nil{
		t.Error(err.Error())
	}
	t.Log("解密效果：")
	t.Log("src bytes:", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))
	t.Log("dst bytes:", hex.EncodeToString(deCryptBytes), "[]byte长度:", len(deCryptBytes))

	t.Log("原字串为：" + srcString)
	t.Log("解密字符为：" + string(deCryptBytes))
	if srcString == string(deCryptBytes) {
		t.Log("TDES加解密验证成功")
	} else {
		t.Error("TDES加解密验证失败")
	}

}
