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

func TestParseUrl(t *testing.T) {
	var testUrl = "https://githun.com/gofunchan/encrypt?a=1&bn=2"
	t.Log("源地址：",testUrl)
	encStr := Base64UrlEncoding(testUrl)
	t.Log("编码后：",encStr)
	dec, err := Base64UrlDecoding(encStr)
	if err != nil {
		t.Error("解码错误：",err)
	}
	t.Log("解码后：",string(dec))

	hostname, path, values, err := SimpleParseUrl(testUrl)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("hostname:",hostname,"path:",path,"values:",values)

}



func TestJsonMarshal(t *testing.T) {
	var testJsonString = `{"name": "fun","age": 20,"gender": "x"}`
	t.Log("源JSON字符串：",testJsonString)
	var rs = make(map[string]interface{})
	err := JsonUnMarshalFromString(testJsonString,&rs)
	if err != nil {
		t.Error("编码错误：",err)
	}
	t.Log("编码后：",rs)
	dec, err := JsonMarshalToString(rs)
	if err != nil {
		t.Error("解码错误：",err)
	}
	t.Log("解码后：",dec)
}

func TestJsonEncode(t *testing.T)  {
	jsonFile := "test.json"
	jsonMap := map[string]interface{}{
		"name": "fun",
		"age": 20,
		"gender": "x",
	}
	t.Log("源数据：",jsonMap)
	err := EncodeToJsonFile(&jsonMap, jsonFile)
	if err != nil {
		t.Fatal("编码错误：",err)
	}
	t.Log("编码成功，生成文件：",jsonFile)

	var rs  = make(map[string]interface{})
	err = DecodeFromJsonFile(jsonFile, &rs)
	if err != nil {
		t.Fatal("解码错误：",err)
	}
	t.Log("解码成功,读取数据：",rs)
}
