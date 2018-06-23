package main

import "fmt"

func fToc(f float64) float64{
	return (f -32 )*5/9
}

func main()  {
	const(
		freezingF = 32.0
		boilingF = 212.0
	)
	fmt.Printf("%g f = %g \n",freezingF,fToc(freezingF))
	fmt.Printf("%g f = %g \n",boilingF,fToc(boilingF))
}