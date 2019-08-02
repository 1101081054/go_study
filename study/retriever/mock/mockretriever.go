package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)
}

func (r *Retriever) Get(url string) string {
	return url
}

func (r *Retriever) Post(url string, form map[string]string) string {
	fmt.Println(r.Contents)
	r.Contents = form["contents"]
	fmt.Println(r.Contents)
	return "ok"
}

func (r Retriever) Set() int {
	return 123
}

func (r Retriever) Life() string {
	return "test"
}