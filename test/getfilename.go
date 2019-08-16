package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	myfolder := `D:\MyPHP\www\fudao_new\Application\Home\View\`
	listFile(myfolder)
}

const domain = "http://fudao.171xue.com"

func listFile(myfolder string) {
	files, _ := ioutil.ReadDir(myfolder)
	for _, file := range files {
		if file.IsDir() {
			listFile(myfolder + "/" + file.Name())
		} else {
			fmt.Println(domain + myfolder + "/" + file.Name())
		}
	}
}