package main

import "fmt"

const limit = 100

func main() {
	sum := (limit * (limit + 1) / 2)
	sum_sqr := sum * sum
	sqr_sum := limit * (limit + 1) * (2*limit + 1) / 6
	fmt.Println(sum_sqr - sqr_sum)
}
