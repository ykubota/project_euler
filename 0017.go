package main

import "fmt"

const limit int = 1000

// ignore space
var str map[int]string = map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
	6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
	11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen",
	16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen", 20: "twenty",
	30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy",
	80: "eighty", 90: "ninety", 1000: "onethousand"}

func main() {
	a := 0
	for i := 1; i <= limit; i++ {
		a += length(i)
	}
	fmt.Println(a)
}

func length(i int) int {
	if str[i] != "" {
		return len(str[i])
	}
	if i < 100 {
		return length(i-i%10) + length(i%10)
	}
	if i%100 == 0 {
		return length(i/100) + len("hundred")
	}
	return length(i/100) + len("hundredand") + length(i%100)
}
