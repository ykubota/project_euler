package main

import "fmt"
import "math"

const limit int = 2000000

func main() {
	// Sieve of Eratosthenes
	c := (limit - 1) / 2
	sieve := make([]bool, c+1)
	max := int(math.Sqrt(float64(limit))-1) / 2

	for i := 1; i <= max; i++ {
		if !sieve[i] {
			for j := 2 * i * (i + 1); j <= c; j += 2*i + 1 {
				sieve[j] = true
			}
		}
	}

	a := 2
	for i := 1; i <= c; i++ {
		if !sieve[i] {
			a += 2*i + 1
		}
	}

	fmt.Println(a)
}
