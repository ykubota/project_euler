package main

import "fmt"

const limit int = 1000

func main() {
	a, n := 0, (limit/3)+1
	r := make([]int, n)
	r[n-2], r[n-1] = 2, 0

	for e := 1; e < limit; e++ {
		for i := 0; i < n-1; i++ {
			// carry
			if r[i+1] > 2*2 {
				r[i] = (2*r[i])%10 + 1
			} else {
				r[i] = (2*r[i])%10 + 0
			}
		}
	}

	for _, v := range r {
		a += v
	}
	fmt.Println(a)
}
