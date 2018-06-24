package main

import (
	"time"
	"runtime"
	"fmt"
)

func main() {
var a [10]int
	for i:=0;i<10;i++{
		go func(i int) {
			for{
				a[i]++
				runtime.Gosched() //让渡协程的控制权
			}
		}(i)
	}


	time.Sleep(time.Microsecond)
	fmt.Println(a)

}


