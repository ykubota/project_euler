package main

import "fmt"

const limit int = 1000

func main() {
	length := 0
	for i := limit; i > 1; i-- {
		if length >= i {
			break
		}

		remain := make([]int, i)
		v, j := 1, 0

		for remain[v] == 0 && v != 0 {
			remain[v] = j
			v *= 10
			v %= i
			j++
		}

		if j-remain[v] > length {
			length = j - remain[v]
		}
	}
	fmt.Println(length + 1)
}
