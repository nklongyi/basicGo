package main

import "fmt"

func main()  {
	p:=person{13,"zhangsan"}
	p.modify()
	fmt.Println(p.string())
}
type person struct {
	age int
	name string
}

func (p person) string() string{
	return p.name
}

func (p *person) modify()  {
	p.name = "lisi"
}