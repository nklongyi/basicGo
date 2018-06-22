package main

import "fmt"

func main()  {
	var s []int
	for i:=0;i<100;i++{
		printSlice(s)
		s=append(s,2*i)
	}
	fmt.Println(s)
	s2 := make([]int,16)
	printSlice(s2)
	s1:=[]int{1,2,3,4}
	copy(s2,s1)

	s2 = append(s2[:3],s2[4:]...)
	fmt.Println(s2)

	fmt.Println("Poping from front")
	v:=s2[0]
	s2 = s2[1:]
	fmt.Println(v)

	fmt.Println("poping from back")
	v1:=s2[len(s2)-1]
	s2=s2[:len(s2)-1]
	fmt.Println(v1)

	fmt.Println(s2)

	m:=make(map[string]int)
	m["key"]=3
	m["set"]=4
	//var m3 map[string]int
	for k,v :=range m{
		fmt.Println(k,v)
	}
	m["cose"]=5
	delete(m,"cose")
	if key,ok :=m["cose"];ok{
		fmt.Println(key)
	}else{
		fmt.Println("no key value")
	}
}

func printSlice(s []int)  {
	fmt.Printf("切片长度为:%d,容量为：%d \n",len(s),cap(s))
}