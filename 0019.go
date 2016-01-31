package main

import (
	"fmt"
	"time"
)

const start, end int = 1901, 2000

func main() {
	a := 0
	for y := start; y <= end; y++ {
		for m := 1; m <= 12; m++ {
			t := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
			if t.Weekday() == 0 {
				a++
			}
		}
	}
	fmt.Println(a)
}
