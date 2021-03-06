package coding

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

/*
序列化rsa公私钥并保存为pem文件
@param privateKey *ecdsa.PrivateKey 使用ecdsa中的GenerateKey方法生成私钥值
@param pemKeyPath string pem公私钥文件的保存路径
*/
func GenECCPemKeyFile(privateKey *ecdsa.PrivateKey, pemKeyPath string) error {
	var privateKeyFile = pemKeyPath + "/EccPrivateKey.pem"
	var publicKeyFile = pemKeyPath + "/EccPublicKey.pem"

	// 一、生成私钥文件
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
	publicKey := privateKey.PublicKey

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

// 获取私钥文件
func ParseECCPrivatePemKey(privateKeyFile string) (privateKey *ecdsa.PrivateKey,err error)  {
	// 将私钥文件中的私钥读出，得到使用pem编码的字符串
	file, err := os.Open(privateKeyFile)
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
	privateKey, err = x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// 解析公钥文件
func ParseECCPublicPemKey(publicKeyFile string) (publicKey *ecdsa.PublicKey,err error)  {
	// 从公钥文件获取钥匙字符串
	file, err := os.Open(publicKeyFile)
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

	publicKey = pubInner.(*ecdsa.PublicKey)

	return publicKey, nil
}