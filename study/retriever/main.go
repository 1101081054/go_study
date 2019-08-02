package main

import (
	"fmt"
	"www/study/retriever/file"
	"www/study/retriever/mock"
)

const url  = "http://www.imooc.com"
func post(poster file.Poster){
	poster.Post(url, map[string]string{
		"name": "ccmouse",
		"course": "golang",
	})
}

func session(poster file.RetrieverPoster) string{
	poster.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return poster.Get(url)

}

//func download(r Retriever) string {
//	return r.Get(url)
//}

func main() {
	var r file.Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	fmt.Println(r.Get(url))
	//fmt.Printf("type: %T; value: %v", r, r)
	//fmt.Println()
	//fmt.Println(download(r))
	//r = real.Riever{}
	//fmt.Printf("type: %T; value: %v", r, r)
	//fmt.Println(r.Get(url))

	fmt.Println("Try a session")

	fmt.Println(session(&retriever))
	var s fmt.Stringer = Stringer1{}
	fmt.Println(s.String())
	//fmt.Println(download(r))
}
