package main

import (
	"fmt"
)

func main() {
	chanal1 := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			chanal1 <- i * i
		}
		close(chanal1)
	}()
	for result := range chanal1 {
		fmt.Println(result)
	}
}
