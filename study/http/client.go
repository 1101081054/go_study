package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil{
		panic(err.Error())
	}

	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", s)
}
