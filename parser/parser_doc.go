package parser

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"replacefile/config"
	"replacefile/utils"

	"github.com/nguyenthenguyen/docx"
)

var KeywordMap map[string]string

func parsedoc(file string) {
	keyword := config.ParserConfig.Keyword
	replace := config.ParserConfig.Replaceword
	if len(replace) > 1 && len(replace) == len(keyword) {
		for i, keyword := range keyword {
			KeywordMap[keyword] = replace[i]
		}
	} else if len(replace) == 1 {
		for _, keyword := range keyword {
			KeywordMap[keyword] = replace[0]
		}
	} else {
		fmt.Printf("keyword和replaceword不对应，请检查配置是否正确 \n")
		os.Exit(3)
	}

	_, _, err := replacefilex(file)
	if err != nil {
		fmt.Printf("write %s failed,err: %s \n", file, err.Error())
	}

	fmt.Printf("Handled %s secceed\n", file)

}

// 2.处理并输出报告
func replacefilex(filePath string) (string, string, error) {
	// 打开一个已有格式的文档，这个是要打开的文档路径。
	r, err := docx.ReadDocxFile(filePath)
	if err != nil {
		fmt.Printf("打开word文档错误，错误信息: %s", err)
	}
	fmt.Println("打开docx完成")
	defer r.Close()

	docx1 := r.Editable()

	for keyword, replace := range KeywordMap {
		if err := docx1.Replace(keyword, replace, -1); err != nil {
			fmt.Println("set docx error:" + err.Error())
		}
	}

	fileDir, fileName := filepath.Split(filePath)
	fileDir += "/replace/"
	fileName += "_replace.docx"
	file := fileDir + fileName

	exist, _ := utils.PathExists(fileDir)
	if !exist {
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			return "", "", errors.New("创建文件夹失败")
		}
	}
	err = docx1.WriteToFile(file)
	if err != nil {
		return "", "", err
	}
	fmt.Println("保存docx文件成功")
	return file, fileName, nil
}
