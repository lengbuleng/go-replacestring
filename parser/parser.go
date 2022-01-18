package parser

import (
	"fmt"
	"path"
	"path/filepath"
	"replacefile/config"
	"replacefile/utils"
	"strings"
)

func Parse() {

	var files []string
	KeywordMap = make(map[string]string)

	files = utils.GetFilesInDir(config.ParserConfig.SrcPath)
	if len(files) == 0 {
		fmt.Printf("文件不存在")
		// os.Exit(3)
	}

	for _, file := range files {
		ext := path.Ext(file)
		_, exit := utils.Find(config.ParserConfig.Ext, ext)
		if exit && strings.Contains(ext, ".docx") {
			parsedoc(file)
		} else {
			fmt.Printf("暂不支持解析扩展名为 %s 的文件 %s\n", ext, filepath.Base(file))
			continue
		}
	}

	// 退出程序并设置退出状态值
	// os.Exit(3)
}
