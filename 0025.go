package main

import (
	"fmt"
	"math/big"
)

const limit int = 1000

func main() {
	n, n1, n2, i := big.NewInt(int64(0)), big.NewInt(int64(2)), big.NewInt(int64(1)), 2
	for len(n.String()) < limit {
		n.Set(n2) // get previous result
		n2.Add(n1, n2)
		n1.Set(n)
		i++
	}
	fmt.Println(i - 1)
}
