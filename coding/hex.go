package coding

import "encoding/hex"

// 返回十六进制字符编码成的字符串
func EncodeHexToString(src []byte) string {
	return hex.EncodeToString(src)
}

// 返回hex编码的字符串hexStr解码的数据，入参字符串需事先被hex编码过
func DecodeHexFromString(hexStr string) ([]byte,error) {
	return hex.DecodeString(hexStr)
}

// 将src的数据编码为EncodedLen(len(src))字节，返回实际写入dst的字节数：EncodedLen(len(src))。
func EncodeHex(dst, src []byte) int {
	return hex.Encode(dst, src)
}

// 将src解码为DecodedLen(len(src))字节，返回实际写入dst的字节数；如遇到非法字符，返回描述错误的error。
func DecodeHex(dst, src []byte) (int, error) {
	return hex.Decode(dst, src)
}