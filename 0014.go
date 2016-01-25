package main

import "fmt"

const limit int = 1000000

func main() {

	// brute + record
	rec := make([]int, limit)
	ans, n, step, max, isRecorded := 0, 0, 0, 0, false

	for i := 1; i < len(rec); i++ {
		isRecorded = false
		n, step = i, 0

		for !isRecorded {
			if n == 1 {
				isRecorded = true
				rec[i] = step
				if step > max {
					max = step
					ans = i
				}
			} else if n < i {
				isRecorded = true
				step = step + rec[n]
				rec[i] = step
				if step > max {
					max = step
					ans = i
				}
			} else {
				if n%2 == 0 {
					n /= 2
				} else {
					n = 3*n + 1
				}
			}
			step++
		}

	}
	fmt.Println(ans)
}
