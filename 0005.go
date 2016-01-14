package main

import "fmt"

const limit = 20

func main() {
	n := []uint{2, 3, 5, 7, 11, 13, 17, 19}
	a := uint(1)
	var t uint
	for i := range n {
		t = n[i]
		for t <= limit {
			t *= n[i]
		}
		t /= n[i]
		a *= t
	}
	fmt.Println(a)
}
