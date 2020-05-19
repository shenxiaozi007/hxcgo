package rootpath

import (
	"os"
	"path/filepath"
)

var RootPath = getRootPath()

//返回根目录
func getRootPath() string {
	var rootPath string

	if rootPath, exist := os.LookupEnv("GOD_SEC"); exist {
		return rootPath
	}

	//或者指定文件变量
	rootPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("根目录地址解析失败：" + err.Error())
	}

	return rootPath
}
