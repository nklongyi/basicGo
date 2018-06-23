package main

import "fmt"

func main()  {

	result := make(chan int)

	go func() {
		sum:= 0
		for i:=0;i<10;i++{
			sum=sum+1
		}
		result<-sum
	}()

	fmt.Println(<-result)
}


