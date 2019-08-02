package main

import "fmt"
import "strings"

func main(){
	tracer := "你Hello，World！"
	camma := strings.Index(tracer, "，")
	pos := strings.Index(tracer[camma:], "Wo")
	fmt.Println(camma, pos, tracer[camma+pos:])
}
