package main

import(
	"fmt"
	"mock"
) 


//go语言 接口由使用者定义(非传统意义上的实现者）
type Req interface {
	Get(url string) string
}


func download(r Req)  string{
	return r.Get("www.baidu.com")
}



func main() {
	var r Retriever
	r = Retriever{"hello world"}
	fmt.Println(download(r))
}


