package main

import "fmt"

var err error



func main()  {

	v :=1
	inc(&v)
	fmt.Println(inc(&v))

}
func inc(p *int)  int{
	*p++
	return *p
}
