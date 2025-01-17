package main

import (
	"fmt"
)
func oddNumbers(){ 
	for i:=1;i<100;i++{
		if i%2 == 1{
			fmt.Println(i)
		}
		
	}
	main()
}
func factorial(){
	var p int
	fmt.Scan(&p)
	t:=1
	for i:=1;i<=p;i++{
		t*=i
	}
	fmt.Println("factorial",t)
	main()
}

