package coding

import "encoding/base64"


// base64字符串编码
func EncodeBase64ToString(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}


// base64字符串解码
func DecodeBase64FromString(data string) ([]byte,error) {
	return base64.StdEncoding.DecodeString(data)
}

