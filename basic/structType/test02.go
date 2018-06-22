package main

//camel case
type Queue []int

func (q *Queue)push(v int)  {
	*q = append(*q,v)
}

func (q *Queue)pop() int{
	head :=(*q)[0]
	*q =(*q)[1:0]
	return head
}

func (q *Queue)IsEmpty() bool {
	return len(*q) == 0
}



