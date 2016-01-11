package main

import "fmt"

const limit = 4000000

func main() {
	sum, a, b, c := 0, 1, 1, 2
	for c < limit {
		sum += c
		a = b + c
		b = c + a
		c = a + b
	}
	fmt.Printf("%d\n", sum)
}
