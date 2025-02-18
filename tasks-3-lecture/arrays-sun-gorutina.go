package main

import "fmt"

func calculateSum(){

}

func goRutina(){
	var n int
	var m int

	fmt.Scan(&n,m)
	arr1 := make([]int, n)
	arr2 := make([]int, m)

	for i:=0;i<n;i++{
		var k int
		fmt.Scan(&k)
		arr1[i] = k
	}
	for i:=0;i<n;i++{
		var k int
		fmt.Scan(&k)
		arr2[i] = k
	}



}