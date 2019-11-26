package coding

import (
	"encoding/base64"
	"net/url"
)
// 用于URL参数传输时的编解码,如果要用在url中，需要使用URLEncoding
func Base64UrlEncoding(src string) string {
	return base64.URLEncoding.EncodeToString([]byte(src))
}

// Base64URL解码
func Base64UrlDecoding(src string) ([]byte,error) {
	return base64.URLEncoding.DecodeString(src)
}

// 简单解析URL地址信息
func SimpleParseUrl(urlStr string) (hostname,path string,values map[string][]string,err error) {
	parseUrl, err := url.Parse(urlStr)
	if err != nil {
		return "","",nil,err
	}
	hostname = parseUrl.Hostname()
	path = parseUrl.Path
	values = parseUrl.Query()
	return
}


