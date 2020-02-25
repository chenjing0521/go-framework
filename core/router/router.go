package router

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func AutoLoadRouter() *http.ServeMux {
	getControllerFile()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello")
	})
	return mux
}

func getControllerFile() {
	workPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	listDir(workPath + "/controller")

}

func listDir(dirName string) {
	fileInfo, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fileInfo {
		fileName := dirName + "/" + file.Name()
		fmt.Println(fileName)
		if file.IsDir() {
			listDir(fileName)
		}
	}
}
