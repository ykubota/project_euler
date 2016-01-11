//https://projecteuler.net/problem=1

package main

import "fmt"

const limit = 1000

func main() {
	i := func(n int) int {
		p := limit / n
		return n * (p * (p + 1)) / 2
	}
	fmt.Printf("%d\n", i(3)+i(5)-i(15))
}
