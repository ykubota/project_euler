package main

import (
	"fmt"
	"math/big"
	"strconv"
)

const limit = 100

func main() {
	f := big.NewInt(int64(limit))
	f.MulRange(1, limit)
	s, a := f.String(), 0

	for i := 0; i < len(s); i++ {
		t, _ := strconv.Atoi(string(s[i]))
		a += t
	}
	fmt.Println(a)
}
