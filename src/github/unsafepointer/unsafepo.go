package main

import (
	"fmt"
	"unsafe"
)

func main() {
	u :=new(user)
	fmt.Println(*u)

	pName:=(*string)(unsafe.Pointer(u))
	*pName = "zhangsan"

	pAge:=(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u))+unsafe.Offsetof(u.age)))
	*pAge=30

	fmt.Println(*u)
	}

type user struct {
	name string
	age int
}



