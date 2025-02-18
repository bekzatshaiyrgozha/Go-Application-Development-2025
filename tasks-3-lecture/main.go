package main

import (
	"fmt"
	"time"
)

func main(){
	go chanel()
	go printMessage("Горутина") 
    printMessage("Негізгі функция")
	go goRutinaSay()
	time.Sleep(time.Second)
	fmt.Println("Hello")
}