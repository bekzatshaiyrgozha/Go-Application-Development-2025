package main

import (
	"fmt"
	"time"
)

func chanel(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Мәлімет ch1-ден"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Мәлімет ch2-ден"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second): // Таймаут қосу
		fmt.Println("Ешқандай канал дайын емес!")
	}
}