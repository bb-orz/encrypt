package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
)


func BaseMd5(src string) string {
	// md5计算
	md5 := md5.New()
	md5Bytes := md5.Sum([]byte(src))
	md5String := hex.EncodeToString(md5Bytes)
	return md5String
}

func BaseSha224(src string) string {
	// sha224
	sha224Bytes := sha256.Sum224([]byte(src))
	sha224String := hex.EncodeToString(sha224Bytes[:])
	return sha224String
}

func BaseSha256(src string) string {
	sha256Bytes := sha256.Sum256([]byte(src))
	sha256String := hex.EncodeToString(sha256Bytes[:])
	return sha256String
}

func BaseSha384(src string) string {
	sha384Bytes := sha512.Sum384([]byte(src))
	sha384String := hex.EncodeToString(sha384Bytes[:])
	return sha384String
}

func BaseSha512(src string) string {
	sha512Bytes := sha512.Sum512([]byte(src))
	sha512String := hex.EncodeToString(sha512Bytes[:])
	return sha512String
}

// 单次输入源的哈希运算
func SingleSha256(src string) string {
	// 哈希运算
	myHash := sha256.Sum256([]byte(src))

	// 十六进制字符串转换输出
	hashString := hex.EncodeToString(myHash[:])

	return hashString
}

// 多次输入源的哈希运算
func MultiSha256(srcs ...string) (string,error) {
	hash := sha256.New()

	// 往哈希运算器里添加数据
	for _, v := range srcs {
		_, err := io.WriteString(hash, v)
		if err != nil {
			return "",err
		}
	}

	// 一般这种方式的参数为nil，因为很少有拼接b哈希运算的情况
	bytes := hash.Sum(nil)
	// 十六进制字符串转换输出
	hashString := hex.EncodeToString(bytes)

	return hashString,nil
}
