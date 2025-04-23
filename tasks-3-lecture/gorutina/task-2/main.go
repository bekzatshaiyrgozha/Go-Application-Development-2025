package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func generate(n int, ch chan int) {
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			ch <- i

		}
	}
	close(ch)

}

func main() {
	ch := make(chan int)
	var n int
	fmt.Scan(&n)

	go generate(n, ch)

	for prime := range ch {
		fmt.Println(prime)
	}

}
