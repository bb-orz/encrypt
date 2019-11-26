package coding

import (
	jsoniter "github.com/json-iterator/go"
	"os"
)

// go数据类型转json字符串
func JsonMarshalToString(data interface{}) (string, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.MarshalToString(&data)
}

// json字符串转go数据类型
func JsonUnMarshalFromString(jsonStr string, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.UnmarshalFromString(jsonStr, v)
}

// 编码go数据类型到json文件
func EncodeToJsonFile(data interface{}, targetJsonFile string) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var targetFile *os.File
	var err error

	// 先打开json文件
	targetFile, err = os.OpenFile(targetJsonFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if os.IsNotExist(err) {
		targetFile, err = os.Create(targetJsonFile)
	}
	defer targetFile.Close()

	encoder := json.NewEncoder(targetFile)
	return encoder.Encode(&data)
	
}

// 解码json文件到go数据类型
func DecodeFromJsonFile(sourceJsonFile string, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	// 先打开json文件
	srcFile, err := os.OpenFile(sourceJsonFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	decoder := json.NewDecoder(srcFile)
	return decoder.Decode(&v)
}
