package main

import (
	"fmt"
)

func fiban(n int, ch chan int) {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}

	close(ch)

}

func main() {
	var n int
	fmt.Scan(&n)

	ch := make(chan int)

	go fiban(n, ch)

	for result := range ch {
		fmt.Println(result)
	}
}
