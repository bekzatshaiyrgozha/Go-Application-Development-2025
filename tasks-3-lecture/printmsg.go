package main

import (
    "fmt"
    "time"
)

func printMessage(message string) {
    for i := 0; i < 5; i++ {
        fmt.Println(message)
        time.Sleep(500 * time.Millisecond) 
    }
}