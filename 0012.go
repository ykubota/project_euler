package main

import (
	"fmt"
	"math"
)

const limit int = 500

func main() {
	a := 1
	for i := 2; i < 100000; i++ {
		a += i
		if divisor(a) >= limit {
			fmt.Println(a)
			break
		}
	}
}

func divisor(n int) int {
	m := 0
	for i := 1; i <= int(math.Ceil(math.Sqrt(float64(n))+1)); i++ {
		if n%i == 0 {
			m += 2
		}
		if i*i == n {
			m -= 1
		}
	}
	return m
}
