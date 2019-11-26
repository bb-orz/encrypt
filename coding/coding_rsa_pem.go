package coding

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

/*
序列化rsa公私钥并保存为pem文件
@param privateKey *rsa.PrivateKey 使用rsa中的GenerateKey方法生成私钥值
@param pemKeyPath string pem公私钥文件的保存路径
*/
func GenRSAPemKeyFile(privateKey *rsa.PrivateKey,pemKeyPath string) error {
	var privateKeyFile = pemKeyPath + "/RSAPrivateKey.pem"
	var publicKeyFile = pemKeyPath + "/RSAPublicKey.pem"

	// 通过x509标准将得到的ras私钥序列化为ASN.1的DER编码字符串
	PKCS1PrivateBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	// 将私钥字符串设置到pem格式块中
	privateBlock := pem.Block{
		Type:  "RSA Private Key",
		Bytes: PKCS1PrivateBytes,
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

	// 生成公钥
	// 从得到的私钥对象中将公钥信息取出
	publicKey := privateKey.PublicKey

	// 通过x509标准将得到的ras公钥序列化为ASN.1的DER编码字符串
	PKCS1PublicBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}

	// 将公钥字符串设置到pem格式块中
	publicBlock := pem.Block{
		Type:  "RSA Public Key",
		Bytes: PKCS1PublicBytes,
	}

	// 通过pem将设置好的数据进行编码，并写入磁盘文件
	publicFile, err := os.Create(publicKeyFile)
	if err != nil {
		return err
	}
	defer publicFile.Close()
	err = pem.Encode(publicFile, &publicBlock)
	if err != nil {
		return err
	}

	return nil
}

// 解析RSA私钥文件
func ParseRSAPrivatePemKey(privateKeyFile string) (privateKey *rsa.PrivateKey,err error)  {
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
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return
}

// 解析RSA公钥文件
func ParseRSAPublicPemKey(publicKeyFile string) (publicKey *rsa.PublicKey,err error) {
	// 将公钥文件中的公钥读出，得到使用pem编码的字符串
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

	publicKey = pubInner.(*rsa.PublicKey)

	return
}