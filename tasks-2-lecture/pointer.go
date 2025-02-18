package main

import "fmt"

func pointers(){
	x:=10
	p:=&x

	fmt.Println("x:",x)
	fmt.Println("p:",p)
	fmt.Println("*p",*p)

}