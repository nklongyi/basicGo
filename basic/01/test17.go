package main

import (
	"strconv"
	"os"
	"bufio"
	"fmt"
)

func main()  {

}

/**
转换二进制
 */
func convertToBin(n int) string{
	reuslt :=""
	for ;n>0 ;n/=2  {
		lsb :=n%2
		reuslt = strconv.Itoa(lsb) + reuslt
	}
	return reuslt
}
//读文本逐行打印
func printFile(filename string)  {
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner :=bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever()  {
	for{
		fmt.Println("hello,world!!!")
	}
}