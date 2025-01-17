package main

import (
	"fmt"
)
func largeNumber(){ // find large  among three numbers
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	if a>b && a>c{
		fmt.Println(a)
	}else if b > a && b > c{
		fmt.Println(b)
	}else{
		fmt.Println(c)
	}
	main()

	
}

func posOrNegOrZero(){
	var k int //find number poss or neg or zero
	fmt.Scan(&k)
	if k > 0{
		fmt.Println("Positive")
	}else if k < 0{
		fmt.Println("Negative")
	}else{
		fmt.Println("Zero")
	}
	main()
}
