package utils

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

//GetFilesInDir 获取文件夹下所有文件
func GetFilesInDir(dir string) (files []string) {
	rd, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("Read dir '%s' failed : %v\n", dir, err)
	}

	for _, fileDir := range rd {
		if !fileDir.IsDir() {
			files = append(files, dir+fileDir.Name())
		}
	}
	return
}

//MoveFile 移动文件
func MoveFile(source, destination string) error {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()
	fi, err := src.Stat()
	if err != nil {
		return err
	}
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	perm := fi.Mode() & os.ModePerm
	dst, err := os.OpenFile(destination, flag, perm)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		dst.Close()
		os.Remove(destination)
		return err
	}
	err = dst.Close()
	if err != nil {
		return err
	}
	err = src.Close()
	if err != nil {
		return err
	}
	err = os.Remove(source)
	if err != nil {
		return err
	}
	return nil
}

func Deal(str string) string {
	// 将字符串转换为rune数组
	srcRunes := []rune(str)
	// 创建一个新的rune数组，用来存放过滤后的数据
	dstRunes := make([]rune, 0, len(srcRunes))
	// 过滤不可见字符，根据unicode对照表的0-32和127都是不可见的字符,保留10换行符
	for _, c := range srcRunes {
		if c >= 0 && c <= 31 && c != 10 {
			continue
		}
		if c == 127 {
			continue
		}
		dstRunes = append(dstRunes, c)
	}

	result := string(dstRunes)
	return result
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Find获取一个切片并在其中查找元素。如果找到它，它将返回它的密钥，否则它将返回-1和一个错误的bool。
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
