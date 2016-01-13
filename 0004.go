package main

import "fmt"

const limit = 999

func isReversible(t int) bool {
	n, r := t, 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return t == r
}

func main() {
	a, t := 0, 0
	for i := limit; i > int(limit/10); i-- {
	L:
		for j := limit; j > int(limit/10); j-- {
			t = i * j
			if !isReversible(t) {
				continue L
			}
			if a == 0 || t > a {
				a = t
			}
		}
	}
	fmt.Println(a)
}
