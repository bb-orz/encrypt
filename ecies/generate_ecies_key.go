package ecies

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"strings"
	"time"

	"crypto/x509"
	"encoding/pem"
	"errors"
	mathRand "math/rand"
	"os"
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
@param randKey 至少36位随机字符串
@param pemKeyPath 生成的公私钥路径
*/
func GenerateKey(randKey, pemKeyPath string) error {

	var err error
	var privateKey *ecdsa.PrivateKey
	var publicKey ecdsa.PublicKey
	var curve elliptic.Curve
	var privateKeyFile = pemKeyPath + "/privateKey.pem"
	var publicKeyFile = pemKeyPath + "/publicKey.pem"

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
	// 通过x509标准将得到的ecc私钥序列化为ASN.1的DER编码字符串
	privateBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}
	// 将私钥字符串设置到pem格式块中
	privateBlock := pem.Block{
		Type:  "ecc private key",
		Bytes: privateBytes,
	}

	// 通过pem将设置好的数据进行编码，并写入磁盘文件
	privateFile, err := os.Create(privateKeyFile)
	if err != nil {
		return err
	}
	defer privateFile.Close()
	err = pem.Encode(privateFile, &privateBlock)
	if err != nil {
		return err
	}

	// 二、生成公钥文件
	// 从得到的私钥对象中将公钥信息取出
	publicKey = privateKey.PublicKey

	// 通过x509标准将得到的ecc公钥序列化为ASN.1的DER编码字符串
	publicBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	// 将公钥字符串设置到pem格式块中
	publicBlock := pem.Block{
		Type:  "ecc public key",
		Bytes: publicBytes,
	}

	// 通过pem将设置好的数据进行编码，并写入磁盘文件
	publicFile, err := os.Create(publicKeyFile)
	if err != nil {
		return err
	}
	err = pem.Encode(publicFile, &publicBlock)
	if err != nil {
		return err
	}

	return nil
}

// 获取私钥文件里的私钥内容
func GetPrivateKeyByPemFile(priKeyFile string) (*ecies.PrivateKey, error) {
	// 将私钥文件中的私钥读出，得到使用pem编码的字符串
	file, err := os.Open(priKeyFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := fileInfo.Size()
	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	// 将得到的字符串解码
	block, _ := pem.Decode(buffer)

	// 使用x509将编码之后的私钥解析出来
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 读取文件的ecdsa私钥转化成ecies私钥
	privateKeyForEcies := ecies.ImportECDSA(privateKey)

	return privateKeyForEcies, nil
}

// 获取公钥文件里的公钥内容
func GetPublicKeyByPemFile(pubKeyFile string) (*ecies.PublicKey, error) {
	var err error
	// 从公钥文件获取钥匙字符串
	file, err := os.Open(pubKeyFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, fileInfo.Size())
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	// 将得到的字符串解码
	block, _ := pem.Decode(buffer)

	// 使用x509将编码之后的公钥解析出来
	pubInner, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey := pubInner.(*ecdsa.PublicKey)

	publicKeyForEcies := ecies.ImportECDSAPublic(publicKey)

	return publicKeyForEcies, nil
}
