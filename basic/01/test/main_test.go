package main

import "testing"

func TestAdd(t *testing.T) {
	sum:=Add(1,2)

	if sum == 3 {
		t.Log("The result is ok")
	}else{
		t.Log("Not good enough!")
	}
}
