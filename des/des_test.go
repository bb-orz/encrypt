package des

import (
	"encoding/hex"
	"testing"
)

const (
	MYKEY = "abcdefgh" // 八字节密匙
	IV    = "aaaabbbb" // CBC模式的初始化向量
)

func TestDES(t *testing.T) {
	t.Log("测试DES对称加解密：")
	srcString := "DES对称加解密测试"
	// 加密
	cryptBytes,err := EncryptDES([]byte(srcString), []byte(MYKEY),[]byte(IV))
	if err != nil{
		t.Error(err.Error())
	}
	t.Log("加密效果：")
	t.Log("src :", hex.EncodeToString([]byte(srcString)), "[]byte长度:", len([]byte(srcString)))
	t.Log("dst :", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))


	//解密
	deCryptBytes, err := DecryptDES(cryptBytes, []byte(MYKEY),[]byte(IV))
	if err != nil{
		t.Error(err.Error())
	}
	t.Log("解密效果：")
	t.Log("src :", hex.EncodeToString(cryptBytes), "[]byte长度:", len(cryptBytes))
	t.Log("dst :", hex.EncodeToString(deCryptBytes), "[]byte长度:", len(deCryptBytes))

	t.Log("原字串为：" + srcString)
	t.Log("解密字符为：" + string(deCryptBytes))
	if srcString == string(deCryptBytes){
		t.Log("DES加解密验证成功")
	}else {
		t.Error("DES加解密验证失败")
	}

}
