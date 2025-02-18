package main

import "fmt"

func calculation(x,y int)(int int){
	return x+y
}

func exercises(){
	var x,y int
	fmt.Scan(&x,&y)

	calculation(x,y)
}