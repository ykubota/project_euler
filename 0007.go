package main

// Crazy slow.

import "fmt"

const limit = 10001

// Prime sieve
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)
	for i := 1; i <= limit; i++ {
		prime := <-ch
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
		if i == limit {
			fmt.Println(prime)
		}
	}
}
