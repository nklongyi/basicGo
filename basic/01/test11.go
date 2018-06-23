package main

import "fmt"

func main() {
	//同步通道
	//ch := make(chan int)
	//go func() {
	//	var sum int = 0
	//	for i := 0; i < 10; i++ {
	//		sum += i
	//	}
	//	ch <- sum
	//}()
	//
	//fmt.Println(<-ch)

	one:=make(chan int)
	two :=make(chan int)
	go func() {
		one<-100
	}()

	go func() {
		v := <-one
		two<- v
	}()

	fmt.Println(<-two)
}
