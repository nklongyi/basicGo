package main

import (
	"fmt"
	"math/cmplx"
	"math"
	"io/ioutil"
)

func main()  {
	euler()
	trangle()
	consts()
	enum()

	const filename  = "abc.txt"
	if content,err := ioutil.ReadFile(filename);err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s \n",content)
	}
    fmt.Println(grade(0),grade(89),grade(-4))
}

func consts()  {
	const fileName  ="hello.xtx"
	const a,b = 3,4

	fmt.Println(fileName,a,b)
}

//枚举
func enum(){

	const(
		cpp=iota
		java
		python
		golang
	)

	const(
		b = 1<<(10*iota)
		kb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,python,golang)
	fmt.Println(b,kb,gb,tb,pb)
}

func euler(){
	c:=3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func trangle()  {
	var a ,b = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(c)
}

func grade(score int) string{
	switch  {
	case score<0 || score >100:
		panic(fmt.Sprintf("wrong score:%d",score))
	case score <60:
		return "F"
	case score<70:
		return "C"
	case score<80:
		return "B"
	case score<=100:
		return "A"
	}
	return "Msg"
}
