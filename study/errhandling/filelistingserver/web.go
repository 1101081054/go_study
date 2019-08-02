package main

import (
	"fmt"
	//"github.com/gpmgo/gopm/modules/log"
	"log"
	"net/http"
	"os"
	"www/study/errhandling/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			r := recover()
			log.Printf("Panic: %v", r)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}()

		err := handler(writer, request)
		if err != nil {
			//log.Warn("Error handling request: %s", err.Error())
			log.Printf("Error handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound //访问的文件或方法不存在
			case os.IsPermission(err):
				code = http.StatusForbidden //无权限访问
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/",  errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
