package main

import (
	"fmt"
	"os"
)

func main() {

	
	//largNumber()
	//posOrNegOrZero()

	//oddNumbers()
	//factorial()

	//StudentAndGrade()
	//books()

	fmt.Println("Ghoose:")
	fmt.Println("1. largNumber()")
	fmt.Println("2. posOrNegOrZero()")
	fmt.Println("3. oddNumbers()")
	fmt.Println("4. factorial()")
	fmt.Println("5. StudentAndGrade()")
	fmt.Println("6. books()")
	fmt.Println("7. arrays()")
	fmt.Println("8. Exit")

	var choose int
	fmt.Scan(&choose)

	switch choose{
	case 1:
		largeNumber()
	case 2:
		posOrNegOrZero()
	case 3:
		oddNumbers()

	case 4:
		factorial()
	case 5:
		StudentAndGrade()
	case 6:
		books()
	case 7:
		arrays()
	case 8:
		os.Exit(0)
	default:
		fmt.Println("Try again")
		main()
	}
	

	



}
