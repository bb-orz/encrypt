package coding

import (
	"encoding/csv"
	"os"
)

func EncodeToCSVFile(data [][]string,targetCSVFile string) error  {
	// 先打开CSV文件
	var targetFile *os.File
	var err error
	targetFile, err = os.OpenFile(targetCSVFile, os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if os.IsNotExist(err) {
		targetFile, err = os.Create(targetCSVFile)
	}
	defer targetFile.Close()

	writer := csv.NewWriter(targetFile)
	err = writer.WriteAll(data)
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}

func DecodeFromCSVFile(sourceCSVFile string)(records [][]string,err error)  {
	// 先打开CSV文件
	srcFile, err := os.OpenFile(sourceCSVFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer srcFile.Close()

	reader := csv.NewReader(srcFile)
	return reader.ReadAll()

}
