package main

import (
	"fmt"
	"math"
)



const (
	Pi = math.Pi
	value = math.Phi
)

func sum (a, b int) int {
	c := a + b
	return c
}

func sum1(a, b int) (int){
	d := a + b
	return d
}

func main() {
	for i, j := 10, 20; i < j; i, j = i+1, j-1 {
		fmt.Printf("i = %v, j = %v\n",i, j)
	}

	var c int
	c = sum1(1, 2)
	fmt.Printf("sum = c = %v", c)
}
