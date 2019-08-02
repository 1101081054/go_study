package queue

type Queue interface {
	Push(element interface{})
	Pop() interface{}
	IsEmpty() bool
}

type Entry []interface{}

func (q *Entry) Push(element interface{}){
	*q = append(*q, element)
}

func (q *Entry) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Entry) IsEmpty() bool {
	return len(*q) == 0
}