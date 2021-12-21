package stack

type Stack struct {
	prev  *Stack
	value int
}

func New() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Pop() int {
	n := this.prev
	this.prev = n.prev
	return n.value
}

func (this *Stack) Push(value int) {
	n := &Stack{this.prev, value}
	this.prev = n
}
