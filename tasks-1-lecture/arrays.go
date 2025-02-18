package main

import "fmt"

func arrays() {
	var lenght int
	fmt.Println("Enter lenght of arr:")
	fmt.Scan(&lenght)

	arr := make([]int, lenght)

	for i := 0; i < lenght; i++ {
		fmt.Scan(&arr[i])
	}

	for _, value := range arr {
		fmt.Println(value)
	}
	main()
}
