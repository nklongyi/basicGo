package main

import "fmt"

func main()  {
	x :=[3]int{1,2,3}
	printArray(x)
}

func printArray(arr [...]int)  {
	for i,v :=range arr{
		fmt.Println(i,v)
	}
}
