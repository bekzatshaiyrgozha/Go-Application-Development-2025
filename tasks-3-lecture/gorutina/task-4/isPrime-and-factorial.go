package main

import (
	"fmt"
	"sync"
)

func factorial(n int, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	k := 1
	for i := 1; i <= n; i++ {
		k *= i

	}
	ch2 <- k

}

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

func generete(n int, ch chan int, ch2 chan int) {
	var wg sync.WaitGroup
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			ch <- i
			wg.Add(1)

			go factorial(i, ch2, &wg)
		}
	}
	close(ch)
	wg.Wait()
	close(ch2)

}

func main() {
	var n int
	fmt.Println("Enter size:")

	fmt.Scan(&n)

	ch := make(chan int)
	ch2 := make(chan int)

	go generete(n, ch, ch2)

	fmt.Println("Prime numbers:")

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("-----------")

	fmt.Println("Prime number factorials:")

	for result2 := range ch2 {
		fmt.Println(result2)
	}

}
