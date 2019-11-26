package coding

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)


// 编码go数据类型到Yaml文件
func EncodeToYamlFile(data interface{}, targetYamlFile string) error {
	// 先打开Yaml文件
	var targetFile *os.File
	var err error
	targetFile, err = os.OpenFile(targetYamlFile, os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if os.IsNotExist(err) {
		targetFile, err = os.Create(targetYamlFile)
	}
	defer targetFile.Close()

	encoder := yaml.NewEncoder(targetFile)
	return encoder.Encode(data)

}

// 解码Yaml文件到go数据类型
func DecodeFromYamlFile(sourceYamlFile string, v interface{}) error {
	data, err := ioutil.ReadFile(sourceYamlFile)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &v)
}

