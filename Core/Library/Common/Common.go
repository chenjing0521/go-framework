package Common

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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
	if isExist, _ := FileExists(path); isExist == false {
		return
	}
	fmt.Println(FileExists(path))

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
}
