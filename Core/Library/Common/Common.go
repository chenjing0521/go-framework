package Common

import (
	"io/ioutil"
	"log"
	"os"
)

var workDir string
var routePath []string

func init() {
	workPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	workDir = workPath
}

//文件夹
func FileExists(path string) (isExist bool, err error) {
	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ListDir(path string) []string {
	osPathSep := string(os.PathSeparator)
	listPath := workDir + osPathSep + path
	if isExist, _ := FileExists(listPath); isExist == false {
		log.Fatal("file is not exist")
	}

	files, err := ioutil.ReadDir(listPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range files {
		childPathName := path + osPathSep + v.Name()
		if v.IsDir() {
			routePath = append(routePath, v.Name())
			ListDir(childPathName)
		}
	}
	return routePath
}
