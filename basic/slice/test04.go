package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s:="yes我爱按你你你"
	for i,ch:=range []rune(s){
		fmt.Printf("%d,%c",i,ch)
	}
	fmt.Println(utf8.RuneCountInString(s))
}
