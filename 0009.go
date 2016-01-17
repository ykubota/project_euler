package main

import "fmt"
import "math"

const limit int = 1000

func main() {
	max_c_square := limit / 2
	max := int(math.Sqrt(float64(max_c_square))) - 1

	var a, b, c, j int

loop:
	for i := 2; i <= max; i++ {
		if max_c_square%i == 0 {
			h := max_c_square / i
			for h%2 == 0 {
				h /= 2
			}
			if i%2 == 1 {
				j = i + 2
			} else {
				j = i + 1
			}

			for j < 2*i && j <= h {
				if h%j == 0 && Gcd(j, i) == 1 {
					d := max_c_square / (j * i)
					k := j - i
					a = d * (i*i - k*k)
					b = 2 * d * i * k
					c = d * (i*i + k*k)
					break loop
				}
				j += 2
			}
		}
	}
	fmt.Println(a * b * c)
}

func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
