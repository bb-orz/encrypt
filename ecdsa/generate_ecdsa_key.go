package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encrypt/coding"
	"time"

	"errors"
	mathRand "math/rand"
	"strings"
)

// 生成指定math/rand字节长度的随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+?=-"
	bytes := []byte(str)
	result := []byte{}

	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/*
生成ECC算法的公钥和私钥文件
根据随机字符串生成，randKey至少36位
*/
func GenerateKey(randKey, pemKeyPath string) error {

	var err error
	var privateKey *ecdsa.PrivateKey
	var curve elliptic.Curve

	// 一、生成私钥文件

	// 根据随机字符串长度设置curve曲线
	length := len(randKey)
	// elliptic包实现了几条覆盖素数有限域的标准椭圆曲线,Curve代表一个短格式的Weierstrass椭圆曲线，其中a=-3
	if length < 224/8 {
		err = errors.New("私钥长度太短，至少为36位！")
		return err
	}

	if length >= 521/8+8 {
		// 长度大于73字节，返回一个实现了P-512的曲线
		curve = elliptic.P521()
	} else if length >= 384/8+8 {
		// 长度大于56字节，返回一个实现了P-384的曲线
		curve = elliptic.P384()
	} else if length >= 256/8+8 {
		// 长度大于40字节，返回一个实现了P-256的曲线
		curve = elliptic.P256()
	} else if length >= 224/8+8 {
		// 长度大于36字节，返回一个实现了P-224的曲线
		curve = elliptic.P224()
	}

	// GenerateKey方法生成私钥
	privateKey, err = ecdsa.GenerateKey(curve, strings.NewReader(randKey))
	if err != nil {
		return err
	}

	// 使用pem编码工具生成公私钥pem文件
	err = coding.GenECCPemKeyFile(privateKey, pemKeyPath)
	if err != nil {
		return err
	}

	return nil
}

