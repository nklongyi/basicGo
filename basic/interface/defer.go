package main

import "fmt"

func trydefer()  {
	defer fmt.Println(3)
	defer fmt.Println(2)

	fmt.Println(1)
}



func main() {
	trydefer()
}
