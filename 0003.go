package main

import "fmt"
import "math"

const limit int64 = 600851475143

func isPrime(p int64) bool {
	var i int64 = 2
	for ; i < p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var i int64 = int64(math.Sqrt(float64(limit)))
	for ; i > 1; i-- {
		if limit%i == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}
