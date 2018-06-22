package main

import "fmt"



func main() {
	arr := [...]int{0,1,2,3,4,5,6,5}
	s:=arr[2:6]
	s1:=arr[:]
	updateSlice(s)
	updateSlice(s1)
	//slice本质是一个view，底层是array数组
	fmt.Println(s)
	fmt.Println(s1)
}

func updateSlice(arr []int)  {
	arr[2]=100
	return
}

