package main

import "fmt"

const limit int = 20

func main() {
	a, n, k := 1, 2*limit, limit
	for i := 1; i <= k; i++ {
		a *= n + 1 - i
		a /= i
	}
	fmt.Println(a)
}
