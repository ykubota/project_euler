package main

import "fmt"

const limit int = 10000

// Sum of propper divisors
func d(a int) int {
	sum, i, self := 1, 2, a
	for a >= i*i && a > 1 {
		if a%i == 0 {
			j := i * i
			a /= i
			for a%i == 0 {
				j *= i
				a /= i
			}
			sum *= j - 1
			sum /= i - 1
		}
		// incremental
		if i == 2 {
			i = 3
		} else {
			i += 2
		}
	}
	if a > 1 {
		sum *= a + 1
	}
	return sum - self
}

func main() {
	result := 0
	for a := 2; a < limit; a++ {
		b := d(a)
		if b > a {
			if d(b) == a {
				result += a + b
			}
		}
	}
	fmt.Println(result)
}
