package ecdsa

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
)

/*
使用ECC算法加密签名，返回签名数据
@param input 加密数据
@param priKeyFile 私钥文件
@param randSign 随机字符串
*/
func CryptSignByEcc(input, priKeyFile, randSign string) (output string, err error) {
	// 获取私钥
	privateKey, err := GetPrivateKeyByPemFile(priKeyFile)
	if err != nil {
		return "", err
	}

	// ecc私钥和随机签字符串数据得到哈希
	r, s, err := ecdsa.Sign(strings.NewReader(randSign), privateKey, []byte(input))
	if err != nil {
		return "", err
	}

	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}

	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}

	// 拼接两个椭圆曲线参数哈希
	var b bytes.Buffer
	writer := gzip.NewWriter(&b)
	defer writer.Close()

	_, err = writer.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}
	writer.Flush()

	return hex.EncodeToString(b.Bytes()), nil
}


/*
使用ECC算法,对密文和明文进行匹配校验
@param srcStr 明文
@param cryptStr  密文
@param publicFile 公钥文件
*/
func VerifyCryptEcc(srcStr, cryptStr, publicFile string) (bool, error) {

	decodeBytes, err := hex.DecodeString(cryptStr)
	if err != nil {
		return false, err
	}

	// 解密签名信息，返回椭圆曲线参数：两个大整数
	rint, sint, err := UnSignCryptEcc(decodeBytes)

	// 获取公钥验证数据
	publicKey, err := GetPublicKeyByPemFile(publicFile)
	if err != nil {
		return false, err
	}
	// 使用公钥、原文、以及签名信息解密后的两个椭圆曲线的大整数参数进行校验
	verify := ecdsa.Verify(publicKey, []byte(srcStr), &rint, &sint)

	return verify, nil
}

// 使用ECC算法解密,返回加密前的椭圆曲线大整数
func UnSignCryptEcc(cryptBytes []byte) (rint, sint big.Int, err error) {
	reader, err := gzip.NewReader(bytes.NewBuffer(cryptBytes))
	if err != nil {
		err = errors.New("decode error," + err.Error())
	}
	defer reader.Close()

	buf := make([]byte, 1024)
	count, err := reader.Read(buf)
	if err != nil {
		err = errors.New("decode read error," + err.Error())
	}

	rs := strings.Split(string(buf[:count]), "+")
	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("decrypt rint fail, " + err.Error())
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("decrypt sint fail, " + err.Error())
		return
	}
	return
}
