package main

import "fmt"

func slices(){
	var n int
	fmt.Scan(&n)
	slice:=make([]int,n)

	for i:=0;i<n;i++{
		slice = append(slice, n)
	}
	for i:=0;i<len(slice);i++{
		fmt.Println(slice[i])
	}
}