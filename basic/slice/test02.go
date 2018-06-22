package main

import "fmt"

func main() {
	arr :=[7]int{0,1,2,3,4,5,6}
	s1 := arr[2:6] //2,3,4,5
	s2 :=s1[3:5] //5
	fmt.Printf("容量是：%d ,内容是：%v,长度是 %d \n",cap(s2),s2,len(s2))
	fmt.Println(s2)

	s2 = append(s2,10)
	s2 = append(s2,10)//append可以扩充容量
	fmt.Printf("append : %v，容量是:%d",s2,cap(s2))


}



