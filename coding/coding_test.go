package coding

import (
	"encoding/hex"
	"testing"
)

func TestBase64(t *testing.T) {
	var base64TestStr = "TestBase64 测试源字符串"
	t.Log("源字符串：",base64TestStr)
	encStr := EncodeBase64ToString(base64TestStr)
	t.Log("编码后：",encStr)
	dec, err := DecodeBase64FromString(encStr)
	if err != nil {
		t.Error("解码错误：",err)
	}
	t.Log("解码后：",string(dec))
}


func TestHex(t *testing.T) {
	var hexTestStr = "Hex 测试源字符串"
	t.Log("源字符串：",hexTestStr)

	t.Log("方式一：")
	encStr := EncodeHexToString([]byte(hexTestStr))
	t.Log("编码后：",encStr)
	dec, err := DecodeHexFromString(encStr)
	if err != nil {
		t.Error("解码错误：",err)
	}
	t.Log("解码后：",string(dec))

	t.Log("方式二：")
	var elen = hex.EncodedLen(len(encStr))
	var encRs = make([]byte,elen)
	i := EncodeHex(encRs, []byte(hexTestStr))
	t.Log("编码字符串结果：",string(encRs),"容器长度：",elen,",实际写入长度：",i)
	var dlen = hex.DecodedLen(len(encRs))
	var decRs = make([]byte,dlen)
	i, err = DecodeHex(decRs, encRs)
	t.Log("解码字符串结果：",string(decRs),"容器长度：",dlen,",实际写入长度：",i)
}
