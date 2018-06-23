package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func eval(a,b int,op string )  (int,error){
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		return a/b,nil
	default:
		return 0,fmt.Errorf("unsupported operation")
	}
}

func add(a,b int) int{
	return a+b
}
/*
函数作为入参,函数式编程
 */
func apply(op func(int ,int ) int,a ,b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName :=runtime.FuncForPC(p).Name()

	fmt.Printf("使用了函数 %s ,输入参数为"+"(%d,%d) \n",opName,a,b)

	return op(a,b)
}

func main()  {
	fmt.Println()
	fmt.Println(apply(add,3,4))
	fmt.Println(sum(1,2,3))


	 x:=3
	var y = &x

	fmt.Println(x)
	*y = 4
	fmt.Println(x)


	var arr [5]int
	arr2 :=[3]int{1,3,5}
	arr3:=[...]int{2,4,323,223,5}
	fmt.Println(arr,arr2,arr3)

	//for i:=0;i<len(arr3);i++{
	//	fmt.Println(arr3[i])
	//}

	for _,v:=range arr3{
		fmt.Println(v)
	}
}

func sum(numbers ...int)  int{
	s :=0
	for i:=range numbers{
		s+=numbers[i]
	}
	return s
}



func addpointer(a,b *int) int {
	v:= *a + *b
	return v
}

func swap(a,b int) (int,int) {
	return b,a
}