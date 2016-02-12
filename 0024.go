package main

import "fmt"

const limit int = 1000000

func p(n int) int {
	if n < 0 {
		return 0
	}
	p := 1
	for i := 1; i <= n; i++ {
		p *= i
	}
	return p
}

func main() {
	s := 10
	r := limit - 1
	perm := make([]int, s)
	num := make([]int, s)
	for i := 0; i < s; i++ {
		perm[i] = i
		num[i] = i
	}

	for i := 1; i < s; i++ {
		j := r / p(s-i)
		r = r % p(s-i)
		fmt.Print(perm[j])
		perm = append(perm[:j], perm[j+1:]...)
	}

	for i := range perm {
		fmt.Print(i)
	}
	fmt.Println()

}
