package main

import "fmt"

const limit int = 28123

func main() {
	divsum := make([]int, limit)
	for i := 1; i < limit; i++ {
		for j := i * 2; j < limit; j += i {
			divsum[j] += i
		}
	}

	var abundants []int
	for i, v := range divsum {
		if v > i {
			abundants = append(abundants, i)
		}
	}

	ss := make([]bool, limit)
	for _, v := range abundants {
		for _, vv := range abundants {
			if v+vv < limit {
				ss[v+vv] = true
			} else {
				break
			}
		}
	}

	a := 0
	for i, v := range ss {
		if !v {
			a += i
		}
	}

	fmt.Println(a)
}
