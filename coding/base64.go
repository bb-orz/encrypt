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

// 用于URL参数传输时的编解码,如果要用在url中，需要使用URLEncoding
func Base64UrlEncoding(src string) string {
	return base64.URLEncoding.EncodeToString([]byte(src))
}

// Base64URL解码
func Base64UrlDecoding(src string) ([]byte,error) {
	return base64.URLEncoding.DecodeString(src)
}
